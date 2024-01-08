package imap

import (
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
)

func getSearchOptions(options ...*imap.SearchOptions) *imap.SearchOptions {
	var opts *imap.SearchOptions

	if len(options) > 0 && options[0] != nil {
		opts = options[0]
	}

	if opts == nil {
		opts = &imap.SearchOptions{
			ReturnMin:   true,
			ReturnMax:   true,
			ReturnAll:   true,
			ReturnCount: true,
		}
	}

	return opts
}

// StreamSearch searches the mailbox for messages that match the given searching criteria.
func StreamSearch(c *imapclient.Client, criteria *imap.SearchCriteria, options ...*imap.SearchOptions) *imapclient.SearchCommand {
	logger.Info().Any("criteria", criteria).Msg("searching message(s)")
	return c.Search(criteria, getSearchOptions(options...))
}

// Search searches the mailbox for messages that match the given searching criteria.
func Search(c *imapclient.Client, criteria *imap.SearchCriteria, options ...*imap.SearchOptions) (*imap.SearchData, error) {
	return StreamSearch(c, criteria, options...).Wait()
}

// StreamUIDSearch searches the mailbox for messages that match the given searching criteria.
func StreamUIDSearch(c *imapclient.Client, criteria *imap.SearchCriteria, options ...*imap.SearchOptions) *imapclient.SearchCommand {
	logger.Info().Any("criteria", criteria).Msg("searching message(s)")
	return c.UIDSearch(criteria, getSearchOptions(options...))
}

// UIDSearch searches the mailbox for messages that match the given searching criteria.
func UIDSearch(c *imapclient.Client, criteria *imap.SearchCriteria, options ...*imap.SearchOptions) (*imap.SearchData, error) {
	return StreamUIDSearch(c, criteria, options...).Wait()
}
