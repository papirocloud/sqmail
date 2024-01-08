package imap

import (
	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/emersion/go-sasl"
)

// Authenticate advises the server that the client wants to initiate a SASL
// authentication mechanism to prove its identity.
// The client may only select a SASL method advertised by the server,
// which is shown in the response to the command CAPABILITY.
func Authenticate(c *imapclient.Client, saslClient sasl.Client) error {
	logger.Info().Msg("authenticating")
	return c.Authenticate(saslClient)
}

// Login advises the server that the client wants to login with a
// username and password to prove its identity.
func Login(c *imapclient.Client, username, password string) error {
	logger.Info().Str("username", username).Msg("logging in")
	return c.Login(username, password).Wait()
}

// Unauthenticate advises the server that the client wants to log out.
func Unauthenticate(c *imapclient.Client) error {
	logger.Info().Msg("unauthenticating")
	return c.Unauthenticate().Wait()
}

// Logout advises the server that the client wants to log out.
func Logout(c *imapclient.Client) error {
	logger.Info().Msg("logging out")
	return c.Logout().Wait()
}
