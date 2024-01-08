package imap

import (
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
)

// Copy copies the specified message(s) to the end of the specified destination mailbox.
func Copy(c *imapclient.Client, seqSet imap.SeqSet, mailbox string) (*imap.CopyData, error) {
	logger.Info().Str("mailbox", mailbox).Msg("copying message(s)")
	return c.Copy(seqSet, mailbox).Wait()
}

// UIDCopy copies the specified message(s) to the end of the specified destination mailbox.
func UIDCopy(c *imapclient.Client, seqSet imap.SeqSet, mailbox string) (*imap.CopyData, error) {
	logger.Info().Str("mailbox", mailbox).Msg("copying message(s)")
	return c.UIDCopy(seqSet, mailbox).Wait()
}
