package email

import (
	"testing"

	"github.com/emersion/go-imap/v2"
)

func TestHasFlag(t *testing.T) {
	m := &Message{
		Flags: []imap.Flag{imap.FlagSeen, imap.FlagDeleted},
	}

	if !m.HasFlag(imap.FlagSeen) {
		t.Errorf("Expected HasFlag to return true for '%s', got false", imap.FlagSeen)
	}

	if m.HasFlag(imap.FlagAnswered) {
		t.Errorf("Expected HasFlag to return false for '%s', got true", imap.FlagAnswered)
	}
}
