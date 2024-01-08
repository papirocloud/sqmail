package sql

import (
	"fmt"

	"github.com/emersion/go-imap/v2"
)

type WhereClause struct {
	Field    *Field
	Operator Operator
	Value    []string
}

// String returns the string representation of the clause.
func (w *WhereClause) String() string {
	return fmt.Sprintf("%s %s %s", w.Field, w.Operator, w.Value)
}

// GetValue returns the first value of the clause.
func (w *WhereClause) GetValue() string {
	if len(w.Value) == 0 {
		return ""
	}

	if w.Field.Name == "headers" {
		return w.Value[1]
	}

	return w.Value[0]
}

func (w *WhereClause) GetKey() string {
	if len(w.Value) == 0 {
		return ""
	}

	return w.Value[0]
}

func (w *WhereClause) ApplyCriteria(criteria *imap.SearchCriteria) {
	if w.Field.ApplyCriteria != nil {
		w.Field.ApplyCriteria(criteria, w)
	}
}
