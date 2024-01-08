package imap

import (
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
)

func getFetchOptions(options ...*imap.FetchOptions) *imap.FetchOptions {
	var opts *imap.FetchOptions

	if len(options) > 0 && options[0] != nil {
		opts = options[0]
	}

	if opts == nil {
		opts = &imap.FetchOptions{
			Envelope:      true,
			Flags:         true,
			InternalDate:  true,
			RFC822Size:    true,
			UID:           true,
			BodyStructure: &imap.FetchItemBodyStructure{},
			BodySection:   []*imap.FetchItemBodySection{{}},
		}
	}

	return opts
}

// StreamFetch retrieves data associated with a message in the mailbox.
func StreamFetch(c *imapclient.Client, seqSet imap.SeqSet, options ...*imap.FetchOptions) *imapclient.FetchCommand {
	logger.Info().Msg("fetching message(s)")
	return c.Fetch(seqSet, getFetchOptions(options...))
}

// Fetch retrieves data associated with a message in the mailbox.
func Fetch(c *imapclient.Client, seqSet imap.SeqSet, options ...*imap.FetchOptions) ([]*imapclient.FetchMessageBuffer, error) {
	return StreamFetch(c, seqSet, options...).Collect()
}

// StreamUIDFetch retrieves data associated with a message in the mailbox.
func StreamUIDFetch(c *imapclient.Client, seqSet imap.SeqSet, options ...*imap.FetchOptions) *imapclient.FetchCommand {
	logger.Info().Msg("fetching message(s)")
	return c.UIDFetch(seqSet, getFetchOptions(options...))
}

// UIDFetch retrieves data associated with a message in the mailbox.
func UIDFetch(c *imapclient.Client, seqSet imap.SeqSet, options ...*imap.FetchOptions) ([]*imapclient.FetchMessageBuffer, error) {
	return StreamUIDFetch(c, seqSet, options...).Collect()
}
