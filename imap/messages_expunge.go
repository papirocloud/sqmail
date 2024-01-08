package imap

import (
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
)

// Expunge permanently removes all messages that have the \Deleted flag set
// from the currently selected mailbox.
func Expunge(c *imapclient.Client) error {
	logger.Info().Msg("expunging mailbox")
	return c.Expunge().Wait()
}

// UIDExpunge permanently removes all messages that have the \Deleted flag set
func UIDExpunge(c *imapclient.Client, seqSet imap.SeqSet) error {
	logger.Info().Msg("expunging mailbox")
	return c.UIDExpunge(seqSet).Wait()
}
