package email

import (
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
)

type EmbeddedFile struct {
	ContentType string
	ContentId   string
	Content     []byte
	Filename    string
}

type Attachment struct {
	ContentType string
	Content     []byte
	Filename    string
}

type Envelope struct {
	Date      time.Time
	Subject   string
	MessageID string
	InReplyTo string
	From      []string
	Sender    []string
	ReplyTo   []string
	To        []string
	Cc        []string
	Bcc       []string
}

type Attendee struct {
	Email  string
	Status ics.ParticipationStatus
}

type Event struct {
	Id          string
	Attendees   []*Attendee
	Description string
	Summary     string
	Link        string
}

type Message struct {
	UID          uint32
	SeqNum       uint32
	Flags        []imap.Flag
	InternalDate time.Time
	RFC822Size   int64

	Mailbox string

	ContentType string

	Raw []byte

	Headers map[string][]string

	RawHeaders []byte
	RawBody    []byte
	RawMime    []byte

	Text string
	HTML string

	Attachments    []*Attachment
	HasAttachments bool

	Embedded  []*EmbeddedFile
	HasEmbeds bool

	Events     []*Event
	HasEvents  bool
	RawInvites [][]byte

	Envelope *Envelope
}

func NewMessage() *Message {
	return &Message{
		Headers:  make(map[string][]string),
		Envelope: &Envelope{},
	}
}

func NewMessageFromIMAP(buffer *imapclient.FetchMessageBuffer, fields map[string]bool) (*Message, error) {
	m := NewMessage()
	err := m.FromIMAP(buffer, fields)
	if err != nil {
		return nil, err
	}

	return m, nil
}
