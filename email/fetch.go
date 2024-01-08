package email

import (
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	sqmailImap "github.com/papirocloud/sqmail/imap"
)

func (m *Message) Fetch(c *imapclient.Client, options ...*imap.FetchOptions) error {
	_, err := sqmailImap.UIDFetch(c, imap.SeqSetNum(m.UID), options...)
	return err
}
