package sql

import (
	"testing"

	"github.com/auxten/postgresql-parser/pkg/sql/sem/tree"
	"github.com/papirocloud/sqmail/email"
)

func TestGetFieldsFromMessage(t *testing.T) {
	m := &email.Message{
		Envelope: &email.Envelope{
			Subject: "Test Subject",
		},
	}

	fields := []*Field{
		{
			Name: "subject",
			GetValue: func(m *email.Message) interface{} {
				return m.Envelope.Subject
			},
		},
	}

	result := GetFieldsFromMessage(m, fields)

	if result["subject"] != "Test Subject" {
		t.Errorf("Expected subject to be 'Test Subject', got '%s'", result["subject"])
	}
}

func TestParseQuery(t *testing.T) {
	query := "SELECT subject FROM mailbox WHERE subject = 'Test' LIMIT 1"
	result, err := ParseQuery(query)

	if err != nil {
		t.Errorf("Expected ParseQuery to return no error, got '%s'", err)
	}

	if result.Clause != "SELECT" {
		t.Errorf("Expected clause to be 'SELECT', got '%s'", result.Clause)
	}

	if len(result.Fields) != 1 || result.Fields[0].Name != "subject" {
		t.Errorf("Expected fields to contain 'subject', got %v", result.Fields)
	}

	if result.From != "mailbox" {
		t.Errorf("Expected from to be 'mailbox', got '%s'", result.From)
	}

	if len(result.Conds) != 1 || result.Conds[0].Field.Name != "subject" || result.Conds[0].Operator != Equals || result.Conds[0].Value[0] != "Test" {
		t.Errorf("Expected conds to contain 'subject = Test', got %v", result.Conds)
	}

	if result.Limit != 1 {
		t.Errorf("Expected limit to be 1, got %d", result.Limit)
	}
}

func TestGetFieldsFromQuery(t *testing.T) {
	query := "SELECT subject FROM mailbox WHERE subject = 'Test' LIMIT 1"
	fields := GetFieldsFromQuery(query)

	if len(fields) != 1 || fields[0].Name != "subject" {
		t.Errorf("Expected fields to contain 'subject', got %v", fields)
	}
}

func TestSanitizeValue(t *testing.T) {
	value := sanitizeValue("'Test'")

	if value != "Test" {
		t.Errorf("Expected value to be 'Test', got '%s'", value)
	}
}

func TestParseLimit(t *testing.T) {
	result := &ParseResult{}
	n := &tree.Limit{
		Count: tree.NewDInt(1),
	}

	err := parseLimit(result, n)

	if err != nil {
		t.Errorf("Expected parseLimit to return no error, got '%s'", err)
	}

	if result.Limit != 1 {
		t.Errorf("Expected limit to be 1, got %d", result.Limit)
	}
}

func TestParseSelect(t *testing.T) {
	result := &ParseResult{}
	n := &tree.Select{}

	err := parseSelect(result, n)

	if err != nil {
		t.Errorf("Expected parseSelect to return no error, got '%s'", err)
	}

	if result.Clause != "SELECT" {
		t.Errorf("Expected clause to be 'SELECT', got '%s'", result.Clause)
	}
}

func TestParseSelectExpr(t *testing.T) {
	result := &ParseResult{}
	n := tree.SelectExpr{
		Expr: tree.NewStrVal("subject"),
	}

	err := parseSelectExpr(result, n)

	if err != nil {
		t.Errorf("Expected parseSelectExpr to return no error, got '%s'", err)
	}

	if len(result.Fields) != 1 || result.Fields[0].Name != "'subject'" {
		t.Errorf("Expected fields to contain 'subject', got %v", result.Fields)
	}
}

func TestParseComparisonExpr(t *testing.T) {
	result := &ParseResult{}
	n := &tree.ComparisonExpr{
		Operator: tree.EQ,
		Left:     tree.NewStrVal("subject"),
		Right:    tree.NewStrVal("Test"),
	}

	err := parseComparisonExpr(result, n)

	if err != nil {
		t.Errorf("Expected parseComparisonExpr to return no error, got '%s'", err)
	}

	if len(result.Conds) != 1 || result.Conds[0].Field.Name != "'subject'" || result.Conds[0].Operator != Equals || result.Conds[0].Value[0] != "Test" {
		t.Errorf("Expected conds to contain 'subject = Test', got %v", result.Conds)
	}
}

func TestParseRangeCond(t *testing.T) {
	result := &ParseResult{}
	n := &tree.RangeCond{
		Left: tree.NewStrVal("subject"),
		From: tree.NewStrVal("Test1"),
		To:   tree.NewStrVal("Test2"),
	}

	err := parseRangeCond(result, n)

	if err != nil {
		t.Errorf("Expected parseRangeCond to return no error, got '%s'", err)
	}

	if len(result.Conds) != 2 || result.Conds[0].Field.Name != "'subject'" || result.Conds[0].Operator != GreaterEq || result.Conds[0].Value[0] != "'Test1'" || result.Conds[1].Field.Name != "'subject'" || result.Conds[1].Operator != LessEq || result.Conds[1].Value[0] != "'Test2'" {
		t.Errorf("Expected conds to contain 'subject >= Test1' and 'subject <= Test2', got %v", result.Conds)
	}
}
