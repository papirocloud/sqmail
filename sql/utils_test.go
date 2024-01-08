package sql

import (
	"testing"

	"github.com/emersion/go-imap/v2"
)

func TestAddBodySection(t *testing.T) {
	opts := &imap.FetchOptions{
		BodySection: []*imap.FetchItemBodySection{
			{
				Specifier: imap.PartSpecifierHeader,
				Peek:      true,
			},
		},
	}

	addBodySection(opts, imap.PartSpecifierText, true)

	found := false
	for _, section := range opts.BodySection {
		if section.Specifier == imap.PartSpecifierText {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected BodySection to contain PartSpecifierText, but it was not found")
	}
}
