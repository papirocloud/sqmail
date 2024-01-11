package email

import (
	"bytes"
	"testing"
	"time"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/mnako/letters"
)

func TestParseBodySection(t *testing.T) {
	buffer := &imapclient.FetchMessageBuffer{
		BodySection: map[*imap.FetchItemBodySection][]byte{
			{}: []byte("Test Body"),
		},
	}

	m := NewMessage()
	m.parseBodySection(buffer)

	if !bytes.Equal(m.Raw, []byte("Test Body")) {
		t.Errorf("Expected raw body to be 'Test Body', got '%s'", string(m.Raw))
	}
}

func TestParseFlags(t *testing.T) {
	buffer := &imapclient.FetchMessageBuffer{
		Flags: []imap.Flag{imap.FlagSeen, imap.FlagDeleted},
	}

	m := NewMessage()
	m.parseFlags(buffer)

	if !m.HasFlag(imap.FlagSeen) {
		t.Errorf("Expected HasFlag to return true for '%s', got false", imap.FlagSeen)
	}

	if !m.HasFlag(imap.FlagDeleted) {
		t.Errorf("Expected HasFlag to return true for '%s', got false", imap.FlagDeleted)
	}
}

func TestParseInternalDate(t *testing.T) {
	buffer := &imapclient.FetchMessageBuffer{
		InternalDate: time.Now(),
	}

	m := NewMessage()
	m.parseInternalDate(buffer)

	if !m.InternalDate.Equal(buffer.InternalDate) {
		t.Errorf("Expected internal date to be '%v', got '%v'", buffer.InternalDate, m.InternalDate)
	}
}

func TestParseSize(t *testing.T) {
	buffer := &imapclient.FetchMessageBuffer{
		RFC822Size: 123,
	}

	m := NewMessage()
	m.parseSize(buffer)

	if m.RFC822Size != buffer.RFC822Size {
		t.Errorf("Expected RFC822 size to be '%d', got '%d'", buffer.RFC822Size, m.RFC822Size)
	}
}

func TestEnsureRaw(t *testing.T) {
	m := &Message{
		RawHeaders: []byte("Test Headers"),
		RawBody:    []byte("Test Body"),
	}

	m.ensureRaw()

	expectedRaw := []byte("Test Headers\r\n\r\nTest Body")
	if !bytes.Equal(m.Raw, expectedRaw) {
		t.Errorf("Expected raw to be '%s', got '%s'", string(expectedRaw), string(m.Raw))
	}
}

func TestParseText(t *testing.T) {
	msg := letters.Email{
		Text:         "Test Text",
		EnrichedText: "Test Enriched Text",
	}

	m := NewMessage()
	m.parseText(msg)

	if m.Text != msg.EnrichedText {
		t.Errorf("Expected text to be '%s', got '%s'", msg.EnrichedText, m.Text)
	}
}

func TestParseHtml(t *testing.T) {
	msg := letters.Email{
		HTML: "Test HTML",
	}

	m := NewMessage()
	m.parseHtml(msg)

	if m.HTML != msg.HTML {
		t.Errorf("Expected HTML to be '%s', got '%s'", msg.HTML, m.HTML)
	}
}

func TestParseHeaders(t *testing.T) {
	msg := letters.Email{
		Headers: letters.Headers{
			ExtraHeaders: map[string][]string{
				"Test-Header": {"Test Value"},
			},
		},
	}

	m := NewMessage()
	m.parseHeaders(msg)

	if m.Headers["Test-Header"][0] != msg.Headers.ExtraHeaders["Test-Header"][0] {
		t.Errorf("Expected header 'Test-Header' to be '%s', got '%s'", msg.Headers.ExtraHeaders["Test-Header"][0], m.Headers["Test-Header"][0])
	}
}
