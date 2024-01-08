package imap

import (
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
)

func Capability(c *imapclient.Client) (imap.CapSet, error) {
	return c.Capability().Wait()
}

func Caps(c *imapclient.Client) imap.CapSet {
	caps := c.Caps()

	var capsStrs []string
	for k := range caps {
		capsStrs = append(capsStrs, string(k))
	}
	logger.Info().Strs("capabilities", capsStrs).Msg("got server capabilities")

	return caps
}
