package imap

import (
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
)

func getStoreOptions(options ...*imap.StoreOptions) *imap.StoreOptions {
	var opts *imap.StoreOptions

	if len(options) > 0 && options[0] != nil {
		opts = options[0]
	}

	if opts == nil {
		opts = &imap.StoreOptions{}
	}

	return opts
}

// StreamStore retrieves data associated with a message in the mailbox.
func StreamStore(c *imapclient.Client, seqSet imap.SeqSet, flags *imap.StoreFlags, options ...*imap.StoreOptions) *imapclient.FetchCommand {
	logger.Info().Msg("storing message(s)")
	return c.Store(seqSet, flags, getStoreOptions(options...))
}

// Store retrieves data associated with a message in the mailbox.
func Store(c *imapclient.Client, seqSet imap.SeqSet, flags *imap.StoreFlags, options ...*imap.StoreOptions) ([]*imapclient.FetchMessageBuffer, error) {
	return StreamStore(c, seqSet, flags, options...).Collect()
}

// StreamUIDStore retrieves data associated with a message in the mailbox.
func StreamUIDStore(c *imapclient.Client, seqSet imap.SeqSet, flags *imap.StoreFlags, options ...*imap.StoreOptions) *imapclient.FetchCommand {
	logger.Info().Msg("storing message(s)")
	return c.UIDStore(seqSet, flags, getStoreOptions(options...))
}

// UIDStore retrieves data associated with a message in the mailbox.
func UIDStore(c *imapclient.Client, seqSet imap.SeqSet, flags *imap.StoreFlags, options ...*imap.StoreOptions) ([]*imapclient.FetchMessageBuffer, error) {
	return StreamUIDStore(c, seqSet, flags, options...).Collect()
}
