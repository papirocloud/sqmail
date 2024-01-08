package sql

import (
	"testing"
)

func TestAddField(t *testing.T) {
	field := func() *Field {
		return &Field{
			Name: "TestField",
		}
	}

	AddField(field)

	if GetField("TestField").Name != "TestField" {
		t.Errorf("Expected field name to be 'TestField', got '%s'", GetField("TestField").Name)
	}
}

func TestListFields(t *testing.T) {
	field := func() *Field {
		return &Field{
			Name: "TestField",
		}
	}

	AddField(field)

	fields := ListFields()

	if len(fields) == 0 {
		t.Errorf("Expected fields to contain at least one field, got %v", fields)
	}
}

func TestGetField(t *testing.T) {
	field := GetField("UnknownField")

	if field.Valid {
		t.Errorf("Expected field to be invalid, got valid")
	}
}

func TestFieldValidate(t *testing.T) {
	field := &Field{
		Name:   "TestField",
		Valid:  false,
		Reason: "Test reason",
	}

	err := field.Validate()

	if err == nil {
		t.Errorf("Expected Validate to return an error, got nil")
	}
}

func TestFieldIsAllowedOperator(t *testing.T) {
	field := &Field{
		Name:             "TestField",
		AllowedOperators: []Operator{Equals},
	}

	if !field.IsAllowedOperator(Equals) {
		t.Errorf("Expected IsAllowedOperator to return true for Equals, got false")
	}

	if field.IsAllowedOperator(Like) {
		t.Errorf("Expected IsAllowedOperator to return false for Like, got true")
	}
}

func TestFieldSetInvalidOperator(t *testing.T) {
	field := &Field{
		Name:             "TestField",
		AllowedOperators: []Operator{Equals},
		Valid:            true,
	}

	field.SetInvalidOperator(Like)

	if field.Valid {
		t.Errorf("Expected field to be invalid, got valid")
	}
}

func TestFieldSetInvalid(t *testing.T) {
	field := &Field{
		Name:  "TestField",
		Valid: true,
	}

	field.SetInvalid("Test reason")

	if field.Valid {
		t.Errorf("Expected field to be invalid, got valid")
	}
}
