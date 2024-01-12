package email

import (
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	sqmailImap "github.com/papirocloud/sqmail/imap"
)

// HasFlag checks if the message has the specified flag.
func (m *Message) HasFlag(flag imap.Flag) bool {
	for _, f := range m.Flags {
		if f == flag {
			return true
		}
	}
	return false
}

// AddFlags adds the specified flags to the message.
func (m *Message) AddFlags(c *imapclient.Client, flags ...imap.Flag) error {
	_, err := sqmailImap.UIDAddFlags(c, imap.UIDSetNum(imap.UID(m.UID)), flags, true)
	if err != nil {
		return err
	}

	m.Flags = append(m.Flags, flags...)

	return nil
}

// DeleteFlags removes the specified flags from the message.
func (m *Message) DeleteFlags(c *imapclient.Client, flags ...imap.Flag) error {
	_, err := sqmailImap.UIDDeleteFlags(c, imap.UIDSetNum(imap.UID(m.UID)), flags, true)
	if err != nil {
		return err
	}

	for _, f := range flags {
		for i, mf := range m.Flags {
			if mf == f {
				m.Flags = append(m.Flags[:i], m.Flags[i+1:]...)
				break
			}
		}
	}

	return nil
}

// SetFlags sets the specified flags for the message.
func (m *Message) SetFlags(c *imapclient.Client, flags ...imap.Flag) error {
	_, err := sqmailImap.UIDSetFlags(c, imap.UIDSetNum(imap.UID(m.UID)), flags, true)
	if err != nil {
		return err
	}

	m.Flags = flags

	return nil
}

// MarkSeen marks the message as seen.
func (m *Message) MarkSeen(c *imapclient.Client) error {
	return m.AddFlags(c, imap.FlagSeen)
}

// MarkUnseen marks the message as unseen.
func (m *Message) MarkUnseen(c *imapclient.Client) error {
	return m.DeleteFlags(c, imap.FlagSeen)
}

// MarkDeleted marks the message to be deleted.
func (m *Message) MarkDeleted(c *imapclient.Client) error {
	return m.AddFlags(c, imap.FlagDeleted)
}

// MarkUndeleted removes the deleted flag from the message.
func (m *Message) MarkUndeleted(c *imapclient.Client) error {
	return m.DeleteFlags(c, imap.FlagDeleted)
}

// MarkFlagged marks the message as flagged / important.
func (m *Message) MarkFlagged(c *imapclient.Client) error {
	return m.AddFlags(c, imap.FlagFlagged)
}

// MarkUnflagged removes the flagged / important flag from the message.
func (m *Message) MarkUnflagged(c *imapclient.Client) error {
	return m.DeleteFlags(c, imap.FlagFlagged)
}

// MarkDraft marks the message as a draft.
func (m *Message) MarkDraft(c *imapclient.Client) error {
	return m.AddFlags(c, imap.FlagDraft)
}

// MarkUndraft removes the draft flag from the message.
func (m *Message) MarkUndraft(c *imapclient.Client) error {
	return m.DeleteFlags(c, imap.FlagDraft)
}
