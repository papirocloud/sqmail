package imap

import (
	"crypto/tls"
	"fmt"
	"mime"
	"net"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/emersion/go-message/charset"
)

func Connect(host string, port int64, tls bool) (*imapclient.Client, error) {
	logger.Info().Str("host", host).Int64("port", port).Bool("tls", tls).Msg("connecting to IMAP server")
	options := &imapclient.Options{
		WordDecoder: &mime.WordDecoder{CharsetReader: charset.Reader},
	}

	if tls {
		return imapclient.DialTLS(
			net.JoinHostPort(
				host, fmt.Sprintf("%d", port)),
			options)
	}

	return imapclient.DialStartTLS(
		net.JoinHostPort(
			host, fmt.Sprintf("%d", port)),
		options)
}

func Close(c *imapclient.Client) error {
	logger.Info().Msg("closing connection")
	return c.Close()
}

// StartTLS instructs the server that the client wishes to encrypt the session now.
// This command may only be issued if the server advertised that STARTTLS is available and supported.
// If a successful key negotiation results after this command is issued,
// then the session is encrypted and IMAP commands continue.
func StartTLS(c *imapclient.Client, config *tls.Config) error {
	logger.Info().Msg("starting TLS")
	return c.StartTLS(config)
}

func State(c *imapclient.Client) imap.ConnState {
	return c.State()
}
