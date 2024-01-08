package email

import (
	"bytes"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/mnako/letters"
)

// FromIMAP converts a raw IMAP message into a Message object.
func (m *Message) FromIMAP(buffer *imapclient.FetchMessageBuffer, fields map[string]bool) error {
	m.UID = buffer.UID
	m.SeqNum = buffer.SeqNum

	var parsers = map[string]func(messageBuffer *imapclient.FetchMessageBuffer){
		"flags":      m.parseFlags,
		"serverdate": m.parseInternalDate,
		"size":       m.parseSize,
		"body":       m.parseBodySection,
	}

	for field, parseFn := range parsers {
		if _, ok := fields[field]; ok || field == "body" {
			parseFn(buffer)
		}
	}

	m.ensureRaw()

	if err := m.FromIMAPEnvelope(buffer.Envelope, fields); err != nil {
		return err
	}

	if !m.shouldKeepProcessing(fields) {
		return nil
	}

	msg, err := letters.ParseEmail(bytes.NewReader(m.Raw))
	if err != nil {
		return err
	}

	var msgParsers = map[string]func(letters.Email){
		"text":        m.parseText,
		"html":        m.parseHtml,
		"headers":     m.parseHeaders,
		"attachments": m.parseAttachments,
		"embedded":    m.parseEmbeds,
	}

	for field, parseFn := range msgParsers {
		if _, ok := fields[field]; ok {
			parseFn(msg)
		}
	}

	return nil
}

func (m *Message) parseBodySection(buffer *imapclient.FetchMessageBuffer) {
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
}

func (m *Message) parseFlags(buffer *imapclient.FetchMessageBuffer) {
	for _, f := range buffer.Flags {
		m.Flags = append(m.Flags, f)
	}
}

func (m *Message) parseInternalDate(buffer *imapclient.FetchMessageBuffer) {
	m.InternalDate = buffer.InternalDate
}

func (m *Message) parseSize(buffer *imapclient.FetchMessageBuffer) {
	m.RFC822Size = buffer.RFC822Size
}

func (m *Message) shouldKeepProcessing(fields map[string]bool) bool {
	for field := range fields {
		switch field {
		case "text", "html", "headers", "attachments", "embedded":
			return true
		}
	}

	return false
}

func (m *Message) ensureRaw() {
	if m.Raw == nil {
		m.Raw = bytes.Join([][]byte{m.rawHeaders, m.rawBody}, []byte("\r\n"))
	}
}

func (m *Message) parseText(msg letters.Email) {
	if msg.EnrichedText != "" {
		m.Text = msg.EnrichedText
	} else {
		m.Text = msg.Text
	}
}

func (m *Message) parseHtml(msg letters.Email) {
	m.HTML = msg.HTML
}

func (m *Message) parseHeaders(msg letters.Email) {
	for k, v := range msg.Headers.ExtraHeaders {
		m.Headers[k] = v
	}
}
