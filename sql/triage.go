package sql

import (
	"github.com/papirocloud/sqmail/email"
)

func TriageMessage(m *email.Message, conds ...*WhereClause) bool {
	var hasTriageField bool
	var results []bool

	for _, c := range conds {
		if c.Field.Triage != nil {
			hasTriageField = true

			results = append(results, c.Field.Triage(m, c))
		}
	}

	if !hasTriageField {
		return true
	}

	for _, r := range results {
		if r {
			return true
		}
	}

	return false
}
