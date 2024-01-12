package sql

import (
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/papirocloud/sqmail/email"
	sqmailImap "github.com/papirocloud/sqmail/imap"
)

func fetchOptionsFromFields(fields []*Field) *imap.FetchOptions {
	var opts imap.FetchOptions

	opts.UID = true
	opts.BodyStructure = &imap.FetchItemBodyStructure{}
	// opts.BodySection = []*imap.FetchItemBodySection{{}}

	for _, field := range fields {
		if field.ApplyFetchOptions != nil {
			field.ApplyFetchOptions(&opts)
		}
	}

	foundHeaders := false
	for k := range opts.BodySection {
		if opts.BodySection[k].Specifier == imap.PartSpecifierHeader {
			foundHeaders = true
		}
	}

	if !foundHeaders && (len(opts.BodySection) != 1 || opts.BodySection[0].Specifier != imap.PartSpecifierNone) {
		addBodySection(&opts, imap.PartSpecifierHeader, true)
	}

	return &opts
}

func Select(c *imapclient.Client, fields []*Field, from string, limit int64, messageCh chan<- *email.Message, clauses ...*WhereClause) error {
	defer close(messageCh)

	var fieldStrs []string
	for _, field := range fields {
		fieldStrs = append(fieldStrs, field.Name)
	}

	logger.Info().Strs("fields", fieldStrs).Str("from", from).Int64("limit", limit).Msg("searching messages")

	criteria := BuildCriteria(clauses...)

	if _, err := sqmailImap.Select(c, from); err != nil {
		return err
	}

	data, err := sqmailImap.UIDSearch(c, criteria)
	if err != nil {
		return err
	}

	fetchCmd := sqmailImap.StreamUIDFetch(c, data.All.(imap.UIDSet), fetchOptionsFromFields(fields))
	// This is hanging for some reason
	// defer func() { _ = fetchCmd.Close() }()

	handleFetchResult(fields, limit, fetchCmd, messageCh, clauses...)

	return nil
}

func handleFetchResult(fields []*Field, limit int64, fetchCmd *imapclient.FetchCommand, messageCh chan<- *email.Message, clauses ...*WhereClause) {
	cnt := 0
	for {
		if limit > 0 && cnt >= int(limit) {
			break
		}

		msg := fetchCmd.Next()
		if msg == nil {
			break
		}

		buf, err := msg.Collect()
		if err != nil {
			logger.Error().Uint32("seqNum", msg.SeqNum).Err(err).Msg("failed to collect message")
			continue
		}

		var fieldsMap = make(map[string]bool)
		for _, field := range fields {
			fieldsMap[field.Name] = true
		}

		message, err := email.NewMessageFromIMAP(buf, fieldsMap)
		if err != nil {
			logger.Error().Uint32("seqNum", msg.SeqNum).Err(err).Msg("failed to parse message")
			continue
		}

		if TriageMessage(message, clauses...) {
			messageCh <- message
			cnt++
		}
	}
}
