package email

import (
	"bytes"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/mnako/letters"
)

func (m *Message) FromIMAPEnvelope(envelope *imap.Envelope, fields map[string]bool) error {
	if envelope != nil {
		if _, ok := fields["date"]; ok {
			m.Envelope.Date = envelope.Date
		}

		if _, ok := fields["subject"]; ok {
			m.Envelope.Subject = envelope.Subject
		}

		if _, ok := fields["messageid"]; ok {
			m.Envelope.MessageID = envelope.MessageID
		}

		if _, ok := fields["inreplyto"]; ok {
			m.Envelope.InReplyTo = envelope.InReplyTo
		}

		if _, ok := fields["from"]; ok {
			for _, from := range envelope.From {
				addr := from.Addr()
				if addr != "" {
					m.Envelope.From = append(m.Envelope.From, addr)
				}
			}
		}

		if _, ok := fields["sender"]; ok {
			for _, sender := range envelope.Sender {
				addr := sender.Addr()
				if addr != "" {
					m.Envelope.Sender = append(m.Envelope.Sender, addr)
				}
			}
		}

		if _, ok := fields["replyto"]; ok {
			for _, replyTo := range envelope.ReplyTo {
				addr := replyTo.Addr()
				if addr != "" {
					m.Envelope.ReplyTo = append(m.Envelope.ReplyTo, addr)
				}
			}
		}

		if _, ok := fields["to"]; ok {
			for _, to := range envelope.To {
				addr := to.Addr()
				if addr != "" {
					m.Envelope.To = append(m.Envelope.To, addr)
				}
			}
		}

		if _, ok := fields["cc"]; ok {
			for _, cc := range envelope.Cc {
				addr := cc.Addr()
				if addr != "" {
					m.Envelope.Cc = append(m.Envelope.Cc, addr)
				}
			}
		}

		if _, ok := fields["bcc"]; ok {
			for _, bcc := range envelope.Bcc {
				addr := bcc.Addr()
				if addr != "" {
					m.Envelope.Bcc = append(m.Envelope.Bcc, addr)
				}
			}
		}
	}

	return nil
}

// FromIMAP converts a raw IMAP message into a Message object.
func (m *Message) FromIMAP(buffer *imapclient.FetchMessageBuffer, fields map[string]bool) error {
	m.UID = buffer.UID
	m.SeqNum = buffer.SeqNum

	if _, ok := fields["flags"]; ok {
		copy(buffer.Flags, m.Flags)
	}

	if _, ok := fields["serverdate"]; ok {
		m.InternalDate = buffer.InternalDate
	}

	if _, ok := fields["size"]; ok {
		m.RFC822Size = buffer.RFC822Size
	}

	for k, v := range buffer.BodySection {
		switch k.Specifier {
		case imap.PartSpecifierHeader:
			m.rawHeaders = v
		case imap.PartSpecifierText:
			m.rawBody = v
		case imap.PartSpecifierNone:
			m.Raw = v
		}
	}

	if m.Raw == nil {
		m.Raw = bytes.Join([][]byte{m.rawHeaders, m.rawBody}, []byte("\r\n"))
	}

	keepProcessing := false
	for field := range fields {
		switch field {
		case "text", "html", "headers", "attachments", "embedded":
			keepProcessing = true
		}
	}

	if err := m.FromIMAPEnvelope(buffer.Envelope, fields); err != nil {
		return err
	}

	if !keepProcessing {
		return nil
	}

	msg, err := letters.ParseEmail(bytes.NewReader(m.Raw))
	if err != nil {
		return err
	}

	if _, ok := fields["text"]; ok {
		if msg.EnrichedText != "" {
			m.Text = msg.EnrichedText
		} else {
			m.Text = msg.Text
		}
	}

	if _, ok := fields["html"]; ok {
		m.HTML = msg.HTML
	}

	if _, ok := fields["headers"]; ok {
		for k, v := range msg.Headers.ExtraHeaders {
			m.Headers[k] = v
		}
	}

	if _, ok := fields["attachments"]; ok {
		handleAttachments(m, msg)
	}

	if _, ok := fields["embedded"]; ok {
		handleEmbeds(m, msg)
	}

	return nil
}
