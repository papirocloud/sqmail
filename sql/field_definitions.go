package sql

import (
	"github.com/emersion/go-imap/v2"
	"github.com/papirocloud/sqmail/email"
)

func init() {
	// uid
	AddField(func() *Field {
		return &Field{
			Name:             "uid",
			AllowedOperators: []Operator{Equals},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.UID
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				opts.UID = true
			},
			ApplyCriteria: buildUidCriteria,
		}
	})

	// seqnum
	AddField(func() *Field {
		return &Field{
			Name:             "seqnum",
			AllowedOperators: []Operator{Equals},
			Valid:            true,
			Selectable:       false,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.SeqNum
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {},
			ApplyCriteria:     buildSeqNumCriteria,
		}
	})

	// serverdate
	AddField(func() *Field {
		return &Field{
			Name:             "serverdate",
			AllowedOperators: []Operator{Equals, Less, Greater, LessEq, GreaterEq},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.InternalDate
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				opts.InternalDate = true
			},
			ApplyCriteria: buildDateCriteria,
		}
	})

	// date
	AddField(func() *Field {
		return &Field{
			Name:             "date",
			AllowedOperators: []Operator{Equals, Less, Greater, LessEq, GreaterEq},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.Envelope.Date
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				opts.Envelope = true
			},
			ApplyCriteria: buildDateCriteria,
		}
	})

	// size
	AddField(func() *Field {
		return &Field{
			Name:             "size",
			AllowedOperators: []Operator{Equals, Less, Greater, LessEq, GreaterEq},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.RFC822Size
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				opts.RFC822Size = true
			},
			ApplyCriteria: buildSizeCriteria,
		}
	})

	// text
	AddField(func() *Field {
		return &Field{
			Name:             "text",
			AllowedOperators: []Operator{Like},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.Text
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				addBodySection(opts, imap.PartSpecifierText, true)
			},
			ApplyCriteria: buildBodyCriteria,
		}
	})

	// html
	AddField(func() *Field {
		return &Field{
			Name:             "html",
			AllowedOperators: []Operator{Like},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.HTML
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				addBodySection(opts, imap.PartSpecifierText, true)
			},
			ApplyCriteria: buildBodyCriteria,
		}
	})

	// mailbox
	AddField(func() *Field {
		return &Field{
			Name:             "mailbox",
			AllowedOperators: []Operator{Equals},
			Valid:            true,
			Selectable:       false,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.Mailbox
			},
		}
	})

	// raw
	AddField(func() *Field {
		return &Field{
			Name:             "raw",
			AllowedOperators: []Operator{},
			Valid:            true,
			Selectable:       true,
			Searchable:       false,
			GetValue: func(m *email.Message) interface{} {
				return m.Raw
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				addBodySection(opts, imap.PartSpecifierHeader, true)
				addBodySection(opts, imap.PartSpecifierText, true)
			},
		}
	})

	// headers
	AddField(func() *Field {
		return &Field{
			Name:             "headers",
			AllowedOperators: []Operator{Equals, Like},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.Headers
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				opts.Envelope = true
				addBodySection(opts, imap.PartSpecifierHeader, true)
			},
			ApplyCriteria: buildHeaderCriteria,
		}
	})

	// subject
	AddField(func() *Field {
		return &Field{
			Name:             "subject",
			AllowedOperators: []Operator{Equals, Like},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.Envelope.Subject
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				opts.Envelope = true
			},
			ApplyCriteria: buildSubjectCriteria,
		}
	})

	// from
	AddField(func() *Field {
		return &Field{
			Name:             "from",
			Aliases:          []string{"from_", "fromAddresses"},
			AllowedOperators: []Operator{Equals, Like},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.Envelope.From
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				opts.Envelope = true
			},
			ApplyCriteria: buildFromCriteria,
		}
	})

	// to
	AddField(func() *Field {
		return &Field{
			Name:             "to",
			Aliases:          []string{"to_", "toAddresses"},
			AllowedOperators: []Operator{Equals, Like},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.Envelope.To
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				opts.Envelope = true
			},
			ApplyCriteria: buildToCriteria,
		}
	})

	// cc
	AddField(func() *Field {
		return &Field{
			Name:             "cc",
			Aliases:          []string{"cc_", "ccAddresses"},
			AllowedOperators: []Operator{Equals, Like},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.Envelope.Cc
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				opts.Envelope = true
			},
			ApplyCriteria: buildCcCriteria,
		}
	})

	// bcc
	AddField(func() *Field {
		return &Field{
			Name:             "bcc",
			Aliases:          []string{"bcc_", "bccAddresses"},
			AllowedOperators: []Operator{Equals, Like},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.Envelope.Bcc
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				opts.Envelope = true
			},
			ApplyCriteria: buildBccCriteria,
		}
	})

	// flags
	AddField(func() *Field {
		return &Field{
			Name:             "flags",
			AllowedOperators: []Operator{Equals, Unequals},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.Flags
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				opts.Flags = true
			},
			ApplyCriteria: buildFlagCriteria,
		}
	})

	// attachments
	AddField(func() *Field {
		return &Field{
			Name:             "attachments",
			AllowedOperators: []Operator{},
			Valid:            true,
			Selectable:       true,
			Searchable:       false,
			GetValue: func(m *email.Message) interface{} {
				return m.Attachments
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				addBodySection(opts, imap.PartSpecifierText, true)
			},
		}
	})

	// has_attachments
	AddField(func() *Field {
		return &Field{
			Name:             "has_attachments",
			AllowedOperators: []Operator{Equals},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.HasAttachments
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				addBodySection(opts, imap.PartSpecifierText, true)
			},
			Triage: func(m *email.Message, clause *WhereClause) bool {
				return m.HasAttachments
			},
		}
	})

	// embedded
	AddField(func() *Field {
		return &Field{
			Name:             "embedded",
			AllowedOperators: []Operator{},
			Valid:            true,
			Selectable:       true,
			Searchable:       false,
			GetValue: func(m *email.Message) interface{} {
				return m.Embedded
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				addBodySection(opts, imap.PartSpecifierText, true)
			},
		}
	})

	// has_embeds
	AddField(func() *Field {
		return &Field{
			Name:             "has_embeds",
			AllowedOperators: []Operator{Equals},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.HasEmbeds
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				addBodySection(opts, imap.PartSpecifierText, true)
			},
			Triage: func(m *email.Message, clause *WhereClause) bool {
				return m.HasEmbeds
			},
		}
	})

	// has_attachment_with_content_type
	AddField(func() *Field {
		return &Field{
			Name:             "has_attachment_with_content_type",
			AllowedOperators: []Operator{Equals},
			Valid:            true,
			Selectable:       false,
			Searchable:       true,
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				addBodySection(opts, imap.PartSpecifierText, true)
			},
			Triage: func(m *email.Message, clause *WhereClause) bool {
				for _, a := range m.Attachments {
					for _, value := range clause.Value {
						if a.ContentType == value {
							return true
						}
					}
				}

				return false
			},
		}
	})

	// events
	AddField(func() *Field {
		return &Field{
			Name:             "events",
			AllowedOperators: []Operator{},
			Valid:            true,
			Selectable:       true,
			Searchable:       false,
			GetValue: func(m *email.Message) interface{} {
				return m.Events
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				addBodySection(opts, imap.PartSpecifierText, true)
			},
		}
	})

	// has_events
	AddField(func() *Field {
		return &Field{
			Name:             "has_events",
			AllowedOperators: []Operator{Equals},
			Valid:            true,
			Selectable:       true,
			Searchable:       true,
			GetValue: func(m *email.Message) interface{} {
				return m.HasEvents
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				addBodySection(opts, imap.PartSpecifierText, true)
			},
			Triage: func(m *email.Message, clause *WhereClause) bool {
				return m.HasEvents
			},
		}
	})

	// raw_invites
	AddField(func() *Field {
		return &Field{
			Name:             "raw_invites",
			AllowedOperators: []Operator{},
			Valid:            true,
			Selectable:       true,
			Searchable:       false,
			GetValue: func(m *email.Message) interface{} {
				return m.RawInvites
			},
			ApplyFetchOptions: func(opts *imap.FetchOptions) {
				addBodySection(opts, imap.PartSpecifierText, true)
			},
		}
	})
}
