package sql

import (
	"strconv"
	"strings"

	"github.com/auxten/postgresql-parser/pkg/sql/parser"
	"github.com/auxten/postgresql-parser/pkg/sql/sem/tree"
	"github.com/auxten/postgresql-parser/pkg/walk"
)

type ParseResult struct {
	Clause  string
	Fields  []*Field
	From    string
	Mailbox string
	Conds   []*WhereClause
	Limit   int64
}

func sanitizeValue(value string) string {
	return strings.Trim(value, "'")
}

func ParseQuery(sql string) (result *ParseResult, err error) {
	result = &ParseResult{}

	w := &walk.AstWalker{
		Fn: func(ctx interface{}, node interface{}) (stop bool) {
			switch n := node.(type) {
			case *tree.Limit:
				i, err := strconv.Atoi(n.Count.String())
				if err != nil {
					return false
				}
				result.Limit = int64(i)
			case *tree.Select:
				result.Clause = "SELECT"
			case tree.SelectExpr:
				field := GetField(n.Expr.String())

				if !field.Selectable {
					field.SetInvalid("field is not selectable")
				}

				result.Fields = append(result.Fields, field)

			case *tree.TableName:
				result.From = n.String()
			case *tree.ComparisonExpr:
				field := GetField(n.Left.String())

				op := Operator(n.Operator.String())

				if !field.IsAllowedOperator(op) {
					field.SetInvalidOperator(op)
				}

				if !field.Searchable {
					field.SetInvalid("field is not searchable")
				}

				var values []string

				switch field.Name {
				case "mailbox":
					result.Mailbox = sanitizeValue(n.Right.String())
				case "headers":
					if t, ok := n.TypedRight().(*tree.Tuple); ok {
						for _, e := range t.Exprs {
							values = append(values, sanitizeValue(e.String()))
						}
					} else {
						field.SetInvalid("invalid value for headers, must be a tuple in the form (key, value)")
					}
				default:
					switch op {
					case In:
						if t, ok := n.TypedRight().(*tree.Tuple); ok {
							for _, e := range t.Exprs {
								values = append(values, sanitizeValue(e.String()))
							}
						}
					default:
						values = append(values, sanitizeValue(n.Right.String()))
					}
				}

				result.Conds = append(result.Conds, &WhereClause{
					Field:    field,
					Operator: op,
					Value:    values,
				})
			case *tree.RangeCond:
				field1 := GetField(n.Left.String())
				field2 := GetField(n.Left.String())

				if !field1.IsAllowedOperator(GreaterEq) {
					field1.SetInvalidOperator(GreaterEq)
				}

				result.Conds = append(result.Conds, &WhereClause{
					Field:    field1,
					Operator: GreaterEq,
					Value:    []string{n.From.String()},
				})

				if !field2.IsAllowedOperator(LessEq) {
					field2.SetInvalidOperator(LessEq)
				}

				result.Conds = append(result.Conds, &WhereClause{
					Field:    field2,
					Operator: LessEq,
					Value:    []string{n.To.String()},
				})
			}
			return false
		},
	}

	var stmts parser.Statements
	stmts, err = parser.Parse(sql)
	if err != nil {
		return
	}

	_, err = w.Walk(stmts, nil)

	for _, field := range result.Fields {
		if err = field.Validate(); err != nil {
			return
		}
	}

	return
}

func GetFieldsFromQuery(query string) []*Field {
	res, err := ParseQuery(query)
	if err != nil {
		return nil
	}

	return res.Fields
}
