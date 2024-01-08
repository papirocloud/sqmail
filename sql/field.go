package sql

import (
	"fmt"

	"github.com/emersion/go-imap/v2"
	"github.com/papirocloud/sqmail/email"
)

type Field struct {
	Name              string
	Aliases           []string
	AllowedOperators  []Operator
	Selectable        bool
	Searchable        bool
	Valid             bool
	Reason            string
	GetValue          func(m *email.Message) interface{}
	ApplyFetchOptions func(options *imap.FetchOptions)
	ApplyCriteria     func(criteria *imap.SearchCriteria, clause *WhereClause)
	Triage            func(m *email.Message, clause *WhereClause) bool
}

var fields = make(map[string]func() *Field)

func AddField(field func() *Field) {
	f := field()
	fields[f.Name] = field
	for k := range f.Aliases {
		fields[f.Aliases[k]] = field
	}
}

func GetField(name string) *Field {
	if field, ok := fields[name]; ok {
		return field()
	}

	return &Field{
		Name:             name,
		AllowedOperators: []Operator{},
		Valid:            false,
		Reason:           "unknown field",
	}
}

func (f *Field) String() string {
	return f.Name
}

func (f *Field) Validate() error {
	if !f.Valid {
		return fmt.Errorf("field %s is invalid: %s", f.Name, f.Reason)
	}

	return nil
}

func (f *Field) IsAllowedOperator(op Operator) bool {
	for _, o := range f.AllowedOperators {
		if o == op {
			return true
		}
	}

	return false
}

func (f *Field) SetInvalidOperator(op Operator) {
	if f.Valid {
		f.SetInvalid(fmt.Sprintf("operator %s is not allowed for field %s", op, f.Name))
	}
}

func (f *Field) SetInvalid(reason string) {
	if f.Valid {
		f.Valid = false
		f.Reason = reason
	}
}

func GetFieldsFromMessage(m *email.Message, fields []*Field) map[string]interface{} {
	result := make(map[string]interface{})

	for _, field := range fields {
		if field.GetValue != nil {
			result[field.Name] = field.GetValue(m)
		}
	}

	return result
}
