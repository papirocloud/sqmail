package sql

import (
	"testing"

	"github.com/emersion/go-imap/v2"
)

func TestBuildUidCriteria(t *testing.T) {
	criteria := &imap.SearchCriteria{}
	clause := &WhereClause{
		Operator: Equals,
		Field:    GetField("uid"),
		Value:    []string{"123"},
	}

	buildUidCriteria(criteria, clause)

	if len(criteria.UID) != 1 || criteria.UID[0].String() != "123" {
		t.Errorf("Expected UID criteria to contain 123, got %v", criteria.UID)
	}
}

func TestBuildSeqNumCriteria(t *testing.T) {
	criteria := &imap.SearchCriteria{}
	clause := &WhereClause{
		Operator: Equals,
		Field:    GetField("seqnum"),
		Value:    []string{"123"},
	}

	buildSeqNumCriteria(criteria, clause)

	if len(criteria.SeqNum) != 1 || criteria.SeqNum[0].String() != "123" {
		t.Errorf("Expected SeqNum criteria to contain 123, got %v", criteria.SeqNum)
	}
}

func TestBuildDateCriteria(t *testing.T) {
	criteria := &imap.SearchCriteria{}
	clause := &WhereClause{
		Operator: Equals,
		Field:    GetField("date"),
		Value:    []string{"2022-01-01"},
	}

	buildDateCriteria(criteria, clause)

	if criteria.Since.IsZero() || criteria.Before.IsZero() {
		t.Errorf("Expected Since and Before criteria to be set, got %v and %v", criteria.Since, criteria.Before)
	}
}

func TestBuildBodyCriteria(t *testing.T) {
	criteria := &imap.SearchCriteria{}
	clause := &WhereClause{
		Operator: Equals,
		Field:    GetField("text"),
		Value:    []string{"test"},
	}

	buildBodyCriteria(criteria, clause)

	if len(criteria.Body) != 1 || criteria.Body[0] != "test" {
		t.Errorf("Expected Body criteria to contain 'test', got %v", criteria.Body)
	}
}

func TestBuildFlagCriteria(t *testing.T) {
	criteria := &imap.SearchCriteria{}
	clause := &WhereClause{
		Operator: Equals,
		Field:    GetField("flags"),
		Value:    []string{"\\Seen"},
	}

	buildFlagCriteria(criteria, clause)

	if len(criteria.Flag) != 1 || criteria.Flag[0] != "\\Seen" {
		t.Errorf("Expected Flag criteria to contain '\\Seen', got %v", criteria.Flag)
	}
}

func TestBuildSizeCriteria(t *testing.T) {
	criteria := &imap.SearchCriteria{}
	clause := &WhereClause{
		Operator: Equals,
		Field:    GetField("size"),
		Value:    []string{"123"},
	}

	buildSizeCriteria(criteria, clause)

	if criteria.Larger != 122 || criteria.Smaller != 124 {
		t.Errorf("Expected Larger and Smaller criteria to be 122 and 124, got %v and %v", criteria.Larger, criteria.Smaller)
	}
}

func TestBuildHeaderCriteria(t *testing.T) {
	criteria := &imap.SearchCriteria{}
	clause := &WhereClause{
		Operator: Equals,
		Field:    GetField("headers"),
		Value:    []string{"foo", "test"},
	}

	buildHeaderCriteria(criteria, clause)

	if len(criteria.Header) != 1 || criteria.Header[0].Key != "foo" || criteria.Header[0].Value != "test" {
		t.Errorf("Expected Header criteria to contain 'subject: test', got %v", criteria.Header)
	}
}

func TestBuildSubjectCriteria(t *testing.T) {
	criteria := &imap.SearchCriteria{}
	clause := &WhereClause{
		Operator: Equals,
		Field:    GetField("subject"),
		Value:    []string{"test"},
	}

	buildSubjectCriteria(criteria, clause)

	if len(criteria.Header) != 1 || criteria.Header[0].Key != "SUBJECT" || criteria.Header[0].Value != "test" {
		t.Errorf("Expected Header criteria to contain 'SUBJECT: test', got %v", criteria.Header)
	}
}

func TestBuildFromCriteria(t *testing.T) {
	criteria := &imap.SearchCriteria{}
	clause := &WhereClause{
		Operator: Equals,
		Field:    GetField("from"),
		Value:    []string{"test@example.com"},
	}

	buildFromCriteria(criteria, clause)

	if len(criteria.Header) != 1 || criteria.Header[0].Key != "FROM" || criteria.Header[0].Value != "test@example.com" {
		t.Errorf("Expected Header criteria to contain 'FROM: test@example.com', got %v", criteria.Header)
	}
}

func TestBuildToCriteria(t *testing.T) {
	criteria := &imap.SearchCriteria{}
	clause := &WhereClause{
		Operator: Equals,
		Field:    GetField("to"),
		Value:    []string{"test@example.com"},
	}

	buildToCriteria(criteria, clause)

	if len(criteria.Header) != 1 || criteria.Header[0].Key != "TO" || criteria.Header[0].Value != "test@example.com" {
		t.Errorf("Expected Header criteria to contain 'TO: test@example.com', got %v", criteria.Header)
	}
}

func TestBuildCcCriteria(t *testing.T) {
	criteria := &imap.SearchCriteria{}
	clause := &WhereClause{
		Operator: Equals,
		Field:    GetField("cc"),
		Value:    []string{"test@example.com"},
	}

	buildCcCriteria(criteria, clause)

	if len(criteria.Header) != 1 || criteria.Header[0].Key != "CC" || criteria.Header[0].Value != "test@example.com" {
		t.Errorf("Expected Header criteria to contain 'CC: test@example.com', got %v", criteria.Header)
	}
}

func TestBuildBccCriteria(t *testing.T) {
	criteria := &imap.SearchCriteria{}
	clause := &WhereClause{
		Operator: Equals,
		Field:    GetField("bcc"),
		Value:    []string{"test@example.com"},
	}

	buildBccCriteria(criteria, clause)

	if len(criteria.Header) != 1 || criteria.Header[0].Key != "BCC" || criteria.Header[0].Value != "test@example.com" {
		t.Errorf("Expected Header criteria to contain 'BCC: test@example.com', got %v", criteria.Header)
	}
}

func TestBuildCriteria(t *testing.T) {
	clauses := []*WhereClause{
		{
			Operator: Equals,
			Field:    GetField("subject"),
			Value:    []string{"test"},
		},
	}

	criteria := BuildCriteria(clauses...)

	if len(criteria.Header) != 1 || criteria.Header[0].Key != "SUBJECT" || criteria.Header[0].Value != "test" {
		t.Errorf("Expected Header criteria to contain 'subject: test', got %v", criteria.Header)
	}
}
