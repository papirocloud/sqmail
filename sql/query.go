package sql

import (
	"fmt"

	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/papirocloud/sqmail/email"
)

func Query(c *imapclient.Client, query string, messageCh chan<- *email.Message) error {
	logger.Info().Str("query", query).Msg("parsing query")

	res, err := ParseQuery(query)
	if err != nil {
		return err
	}

	if res.Mailbox == "" {
		res.Mailbox = "INBOX"
	}

	switch res.From {
	case "emails":
		switch res.Clause {
		case "SELECT":
			return Select(c, res.Fields, res.Mailbox, res.Limit, messageCh, res.Conds...)
		default:
			return fmt.Errorf("invalid clause %s", res.Clause)
		}
	default:
		return fmt.Errorf("invalid source %s", res.From)
	}
}
