package sql

import (
	"github.com/emersion/go-imap/v2"
)

func addBodySection(opts *imap.FetchOptions, section imap.PartSpecifier, peek bool) {
	found := false
	for k := range opts.BodySection {
		if opts.BodySection[k].Specifier == section {
			found = true
			break
		}
	}

	if !found {
		opts.BodySection = append(opts.BodySection, &imap.FetchItemBodySection{
			Specifier: section,
			Peek:      peek,
		})
	}
}
