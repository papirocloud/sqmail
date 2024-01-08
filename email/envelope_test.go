package email

import (
	"testing"
	"time"

	"github.com/emersion/go-imap/v2"
)

func TestParseEnvelopeDate(t *testing.T) {
	envelope := &imap.Envelope{
		Date: time.Now(),
	}

	m := NewMessage()
	m.parseEnvelopeDate(envelope)

	if !m.Envelope.Date.Equal(envelope.Date) {
		t.Errorf("Expected date to be '%v', got '%v'", envelope.Date, m.Envelope.Date)
	}
}

func TestParseEnvelopeSubject(t *testing.T) {
	envelope := &imap.Envelope{
		Subject: "Test Subject",
	}

	m := NewMessage()
	m.parseEnvelopeSubject(envelope)

	if m.Envelope.Subject != envelope.Subject {
		t.Errorf("Expected subject to be '%s', got '%s'", envelope.Subject, m.Envelope.Subject)
	}
}

func TestParseEnvelopeMessageID(t *testing.T) {
	envelope := &imap.Envelope{
		MessageID: "123",
	}

	m := NewMessage()
	m.parseEnvelopeMessageID(envelope)

	if m.Envelope.MessageID != envelope.MessageID {
		t.Errorf("Expected message ID to be '%s', got '%s'", envelope.MessageID, m.Envelope.MessageID)
	}
}

func TestParseEnvelopeInReplyTo(t *testing.T) {
	envelope := &imap.Envelope{
		InReplyTo: "456",
	}

	m := NewMessage()
	m.parseEnvelopeInReplyTo(envelope)

	if m.Envelope.InReplyTo != envelope.InReplyTo {
		t.Errorf("Expected in-reply-to to be '%s', got '%s'", envelope.InReplyTo, m.Envelope.InReplyTo)
	}
}

func TestParseEnvelopeFrom(t *testing.T) {
	envelope := &imap.Envelope{
		From: []imap.Address{{Mailbox: "test", Host: "example.com"}},
	}

	m := NewMessage()
	m.parseEnvelopeFrom(envelope)

	expectedAddress := envelope.From[0].Addr()
	if m.Envelope.From[0] != expectedAddress {
		t.Errorf("Expected from address to be '%s', got '%s'", expectedAddress, m.Envelope.From[0])
	}
}

func TestParseEnvelopeSender(t *testing.T) {
	envelope := &imap.Envelope{
		Sender: []imap.Address{{Mailbox: "test", Host: "example.com"}},
	}

	m := NewMessage()
	m.parseEnvelopeSender(envelope)

	expectedAddress := envelope.Sender[0].Addr()
	if m.Envelope.Sender[0] != expectedAddress {
		t.Errorf("Expected sender address to be '%s', got '%s'", expectedAddress, m.Envelope.Sender[0])
	}
}

func TestParseEnvelopeReplyTo(t *testing.T) {
	envelope := &imap.Envelope{
		ReplyTo: []imap.Address{{Mailbox: "test", Host: "example.com"}},
	}

	m := NewMessage()
	m.parseEnvelopeReplyTo(envelope)

	expectedAddress := envelope.ReplyTo[0].Addr()
	if m.Envelope.ReplyTo[0] != expectedAddress {
		t.Errorf("Expected reply-to address to be '%s', got '%s'", expectedAddress, m.Envelope.ReplyTo[0])
	}
}

func TestParseEnvelopeTo(t *testing.T) {
	envelope := &imap.Envelope{
		To: []imap.Address{{Mailbox: "test", Host: "example.com"}},
	}

	m := NewMessage()
	m.parseEnvelopeTo(envelope)

	expectedAddress := envelope.To[0].Addr()
	if m.Envelope.To[0] != expectedAddress {
		t.Errorf("Expected to address to be '%s', got '%s'", expectedAddress, m.Envelope.To[0])
	}
}

func TestParseEnvelopeCc(t *testing.T) {
	envelope := &imap.Envelope{
		Cc: []imap.Address{{Mailbox: "test", Host: "example.com"}},
	}

	m := NewMessage()
	m.parseEnvelopeCc(envelope)

	expectedAddress := envelope.Cc[0].Addr()
	if m.Envelope.Cc[0] != expectedAddress {
		t.Errorf("Expected cc address to be '%s', got '%s'", expectedAddress, m.Envelope.Cc[0])
	}
}

func TestParseEnvelopeBcc(t *testing.T) {
	envelope := &imap.Envelope{
		Bcc: []imap.Address{{Mailbox: "test", Host: "example.com"}},
	}

	m := NewMessage()
	m.parseEnvelopeBcc(envelope)

	expectedAddress := envelope.Bcc[0].Addr()
	if m.Envelope.Bcc[0] != expectedAddress {
		t.Errorf("Expected bcc address to be '%s', got '%s'", expectedAddress, m.Envelope.Bcc[0])
	}
}
