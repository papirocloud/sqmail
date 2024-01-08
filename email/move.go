package email

import (
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	sqmailImap "github.com/papirocloud/sqmail/imap"
)

// Move copies the specified message(s) to the end of the specified destination mailbox.
func (m *Message) Move(c *imapclient.Client, mailbox string) error {
	_, err := sqmailImap.UIDMove(c, imap.SeqSetNum(m.UID), mailbox)
	return err
}
