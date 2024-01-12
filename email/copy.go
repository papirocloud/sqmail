package email

import (
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	sqmailImap "github.com/papirocloud/sqmail/imap"
)

// Copy copies the specified message(s) to the end of the specified destination mailbox.
func (m *Message) Copy(c *imapclient.Client, mailbox string) error {
	_, err := sqmailImap.UIDCopy(c, imap.UIDSetNum(imap.UID(m.UID)), mailbox)
	return err
}
