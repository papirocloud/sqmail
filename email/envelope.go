package email

import (
	"github.com/emersion/go-imap/v2"
)

func (m *Message) FromIMAPEnvelope(envelope *imap.Envelope, fields map[string]bool) error {
	if envelope == nil {
		return nil
	}

	var parsers = map[string]func(*imap.Envelope){
		"date":      m.parseEnvelopeDate,
		"subject":   m.parseEnvelopeSubject,
		"messageid": m.parseEnvelopeMessageID,
		"inreplyto": m.parseEnvelopeInReplyTo,
		"from":      m.parseEnvelopeFrom,
		"sender":    m.parseEnvelopeSender,
		"replyto":   m.parseEnvelopeReplyTo,
		"to":        m.parseEnvelopeTo,
		"cc":        m.parseEnvelopeCc,
		"bcc":       m.parseEnvelopeBcc,
	}

	for field, parseFn := range parsers {
		if _, ok := fields[field]; ok {
			parseFn(envelope)
		}
	}

	return nil
}

func (m *Message) parseEnvelopeDate(envelope *imap.Envelope) {
	m.Envelope.Date = envelope.Date
}

func (m *Message) parseEnvelopeSubject(envelope *imap.Envelope) {
	m.Envelope.Subject = envelope.Subject
}

func (m *Message) parseEnvelopeMessageID(envelope *imap.Envelope) {
	m.Envelope.MessageID = envelope.MessageID
}

func (m *Message) parseEnvelopeInReplyTo(envelope *imap.Envelope) {
	m.Envelope.InReplyTo = envelope.InReplyTo
}

func (m *Message) parseEnvelopeFrom(envelope *imap.Envelope) {
	for _, from := range envelope.From {
		addr := from.Addr()
		if addr != "" {
			m.Envelope.From = append(m.Envelope.From, addr)
		}
	}
}

func (m *Message) parseEnvelopeSender(envelope *imap.Envelope) {
	for _, sender := range envelope.Sender {
		addr := sender.Addr()
		if addr != "" {
			m.Envelope.Sender = append(m.Envelope.Sender, addr)
		}
	}
}

func (m *Message) parseEnvelopeReplyTo(envelope *imap.Envelope) {
	for _, replyTo := range envelope.ReplyTo {
		addr := replyTo.Addr()
		if addr != "" {
			m.Envelope.ReplyTo = append(m.Envelope.ReplyTo, addr)
		}
	}
}

func (m *Message) parseEnvelopeTo(envelope *imap.Envelope) {
	for _, to := range envelope.To {
		addr := to.Addr()
		if addr != "" {
			m.Envelope.To = append(m.Envelope.To, addr)
		}
	}
}

func (m *Message) parseEnvelopeCc(envelope *imap.Envelope) {
	for _, cc := range envelope.Cc {
		addr := cc.Addr()
		if addr != "" {
			m.Envelope.Cc = append(m.Envelope.Cc, addr)
		}
	}
}

func (m *Message) parseEnvelopeBcc(envelope *imap.Envelope) {
	for _, bcc := range envelope.Bcc {
		addr := bcc.Addr()
		if addr != "" {
			m.Envelope.Bcc = append(m.Envelope.Bcc, addr)
		}
	}
}
