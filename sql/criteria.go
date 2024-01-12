package sql

import (
	"strconv"

	"github.com/araddon/dateparse"
	"github.com/emersion/go-imap/v2"
)

func buildUidCriteria(criteria *imap.SearchCriteria, clause *WhereClause) {
	switch clause.Operator {
	case Equals:
		i, err := strconv.Atoi(clause.GetValue())
		if err != nil {
			return
		}

		criteria.UID = append(criteria.UID, imap.UIDSetNum(imap.UID(i)))
	default:
		return
	}
}

func buildSeqNumCriteria(criteria *imap.SearchCriteria, clause *WhereClause) {
	switch clause.Operator {
	case Equals:
		i, err := strconv.Atoi(clause.GetValue())
		if err != nil {
			return
		}

		criteria.SeqNum = append(criteria.SeqNum, imap.SeqSetNum(uint32(i)))
	default:
		return
	}
}

func buildDateCriteria(criteria *imap.SearchCriteria, clause *WhereClause) {
	value, err := dateparse.ParseStrict(clause.GetValue())
	if err != nil {
		return
	}

	switch clause.Operator {
	case Equals:
		criteria.Since = value.AddDate(0, 0, -1)
		criteria.Before = value.AddDate(0, 0, 1)
	case Greater:
		criteria.Since = value
	case Less:
		criteria.Before = value
	case GreaterEq:
		criteria.Since = value.AddDate(0, 0, -1)
	case LessEq:
		criteria.Before = value.AddDate(0, 0, 1)
	default:
		return
	}
}

func buildBodyCriteria(criteria *imap.SearchCriteria, clause *WhereClause) {
	switch clause.Operator {
	case Equals, Like:
		criteria.Body = append(criteria.Body, clause.GetValue())
	default:
		return
	}
}

func buildFlagCriteria(criteria *imap.SearchCriteria, clause *WhereClause) {
	switch clause.Operator {
	case Equals:
		criteria.Flag = append(criteria.Flag, imap.Flag(clause.GetValue()))
	case Unequals:
		criteria.NotFlag = append(criteria.NotFlag, imap.Flag(clause.GetValue()))
	default:
		return
	}
}

func buildSizeCriteria(criteria *imap.SearchCriteria, clause *WhereClause) {
	switch clause.Operator {
	case Equals:
		i, err := strconv.Atoi(clause.GetValue())
		if err != nil {
			return
		}

		criteria.Larger = int64(i) - 1
		criteria.Smaller = int64(i) + 1
	case Greater:
		i, err := strconv.Atoi(clause.GetValue())
		if err != nil {
			return
		}

		criteria.Larger = int64(i)
	case Less:
		i, err := strconv.Atoi(clause.GetValue())
		if err != nil {
			return
		}

		criteria.Smaller = int64(i)
	case GreaterEq:
		i, err := strconv.Atoi(clause.GetValue())
		if err != nil {
			return
		}

		criteria.Larger = int64(i) - 1
	case LessEq:
		i, err := strconv.Atoi(clause.GetValue())
		if err != nil {
			return
		}

		criteria.Smaller = int64(i) + 1
	default:
		return
	}
}

func addHeaderCriteria(criteria *imap.SearchCriteria, key, value string) {
	criteria.Header = append(criteria.Header, imap.SearchCriteriaHeaderField{
		Key:   key,
		Value: value,
	})
}

func buildHeaderCriteria(criteria *imap.SearchCriteria, clause *WhereClause) {
	switch clause.Operator {
	case Equals, Like:
		addHeaderCriteria(criteria, clause.GetKey(), clause.GetValue())
	default:
		return
	}
}

func buildSubjectCriteria(criteria *imap.SearchCriteria, clause *WhereClause) {
	switch clause.Operator {
	case Equals, Like:
		addHeaderCriteria(criteria, "SUBJECT", clause.GetValue())
	default:
		return
	}
}

func buildFromCriteria(criteria *imap.SearchCriteria, clause *WhereClause) {
	switch clause.Operator {
	case Equals, Like:
		addHeaderCriteria(criteria, "FROM", clause.GetValue())
	default:
		return
	}
}

func buildToCriteria(criteria *imap.SearchCriteria, clause *WhereClause) {
	switch clause.Operator {
	case Equals, Like:
		addHeaderCriteria(criteria, "TO", clause.GetValue())
	default:
		return
	}
}

func buildCcCriteria(criteria *imap.SearchCriteria, clause *WhereClause) {
	switch clause.Operator {
	case Equals, Like:
		addHeaderCriteria(criteria, "CC", clause.GetValue())
	default:
		return
	}
}

func buildBccCriteria(criteria *imap.SearchCriteria, clause *WhereClause) {
	switch clause.Operator {
	case Equals, Like:
		addHeaderCriteria(criteria, "BCC", clause.GetValue())
	default:
		return
	}
}

func BuildCriteria(clauses ...*WhereClause) *imap.SearchCriteria {
	var criteria imap.SearchCriteria

	for k := range clauses {
		clauses[k].ApplyCriteria(&criteria)
	}

	return &criteria
}
