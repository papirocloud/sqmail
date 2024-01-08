package imap

import (
	"strings"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
)

// Create creates a new mailbox or folder.
func Create(c *imapclient.Client, mailbox string) error {
	logger.Info().Str("mailbox", mailbox).Msg("creating mailbox")
	return c.Create(mailbox, &imap.CreateOptions{}).Wait()
}

// Delete deleted a named mailbor or folder.
func Delete(c *imapclient.Client, mailbox string) error {
	logger.Info().Str("mailbox", mailbox).Msg("deleting mailbox")
	return c.Delete(mailbox).Wait()
}

// Examine does the exact same thing as SELECT, except that it selects the folder in read-only mode,
// meaning that no changes can be effected on the folder.
func Examine(c *imapclient.Client, mailbox string) (*imap.SelectData, error) {
	logger.Info().Str("mailbox", mailbox).Msg("examining mailbox")
	return c.Select(mailbox, &imap.SelectOptions{ReadOnly: true}).Wait()
}

// Select instructs the server that the client now wishes to select a particular mailbox or folder,
// and any commands that relate to a folder should assume this folder as the target of that command.
func Select(c *imapclient.Client, mailbox string) (*imap.SelectData, error) {
	if strings.EqualFold(mailbox, "anywhere") {
		mailbox = "Archive"
	}

	logger.Info().Str("mailbox", mailbox).Msg("selecting mailbox")

	return c.Select(mailbox, &imap.SelectOptions{ReadOnly: false}).Wait()
}

// Unselect does the exact same thing as CLOSE, except that it does not expunge deleted messages.
func Unselect(c *imapclient.Client) error {
	logger.Info().Msg("unselecting mailbox")
	return c.Unselect().Wait()
}

// Rename renames a mailbox or folder from one name to a new name.
func Rename(c *imapclient.Client, existingName, newName string) error {
	logger.Info().Str("existingName", existingName).Str("newName", newName).Msg("renaming mailbox")
	return c.Rename(existingName, newName).Wait()
}

// Subscribe adds the specified mailbox name to the server's set of "active" or "subscribed"
// mailboxes as returned by the LSUB command.
func Subscribe(c *imapclient.Client, mailbox string) error {
	logger.Info().Str("mailbox", mailbox).Msg("subscribing to mailbox")
	return c.Subscribe(mailbox).Wait()
}

// Unsubscribe removes the specified mailbox name from the server's set of "active" or "subscribed"
// mailboxes as returned by the LSUB command.
func Unsubscribe(c *imapclient.Client, mailbox string) error {
	logger.Info().Str("mailbox", mailbox).Msg("unsubscribing from mailbox")
	return c.Unsubscribe(mailbox).Wait()
}

// List lists all the mailboxes and folders present within the server's namespace.
func List(c *imapclient.Client, referenceName, mailboxPattern string) ([]*imap.ListData, error) {
	logger.Info().Str("referenceName", referenceName).Str("mailboxPattern", mailboxPattern).Msg("listing mailboxes")
	return c.List(referenceName, mailboxPattern, &imap.ListOptions{
		ReturnStatus: &imap.StatusOptions{
			NumMessages: true,
			NumUnseen:   true,
		},
	}).Collect()
}

// LSub lists all the mailboxes and folders present within the server's namespace.
func LSub(c *imapclient.Client, referenceName, mailboxPattern string) ([]*imap.ListData, error) {
	logger.Info().Str("referenceName", referenceName).Str("mailboxPattern", mailboxPattern).Msg("listing subscribed mailboxes")
	return c.List(referenceName, mailboxPattern, &imap.ListOptions{
		ReturnStatus: &imap.StatusOptions{
			NumMessages: true,
			NumUnseen:   true,
		},
		SelectSubscribed: true,
	}).Collect()
}

// Status requests the status of the indicated mailbox.
func Status(c *imapclient.Client, mailbox string, options *imap.StatusOptions) (*imap.StatusData, error) {
	logger.Info().Str("mailbox", mailbox).Msg("requesting mailbox status")
	return c.Status(mailbox, options).Wait()
}
