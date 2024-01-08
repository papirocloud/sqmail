package imap

import (
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
)

// Move moves the specified message(s) to the end of the specified destination mailbox.
func Move(c *imapclient.Client, seqSet imap.SeqSet, mailbox string) (*imapclient.MoveData, error) {
	logger.Info().Str("mailbox", mailbox).Msg("moving message(s)")
	return c.Move(seqSet, mailbox).Wait()
}

// UIDMove moves the specified message(s) to the end of the specified destination mailbox.
func UIDMove(c *imapclient.Client, seqSet imap.SeqSet, mailbox string) (*imapclient.MoveData, error) {
	logger.Info().Str("mailbox", mailbox).Msg("moving message(s)")
	return c.UIDMove(seqSet, mailbox).Wait()
}
