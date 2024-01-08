package email

import (
	"bytes"
	"encoding/base64"
	"io"

	ics "github.com/arran4/golang-ical"
)

func handleInvitation(m *Message, r io.Reader) error {
	m.HasEvents = true

	var raw = bytes.NewBuffer(nil)

	enc := base64.NewEncoder(base64.StdEncoding, raw)

	cal, err := ics.ParseCalendar(io.TeeReader(r, enc))
	if err != nil {
		return err
	}
	_ = enc.Close()

	m.RawInvites = append(m.RawInvites, raw.Bytes())

	for _, e := range cal.Events() {
		skip := false
		for _, ee := range m.Events {
			if ee.Id == e.Id() {
				skip = true
				break
			}
		}

		if skip {
			continue
		}

		if err := handleEvent(m, e); err != nil {
			return err
		}
	}

	return nil
}

func handleEvent(m *Message, e *ics.VEvent) error {
	var ev Event

	ev.Id = e.Id()

	for _, a := range e.Attendees() {
		handleAttendee(&ev, a)
	}

	for _, v := range e.Properties {
		switch v.IANAToken {
		case "DESCRIPTION":
			ev.Description = v.Value
		case "SUMMARY":
			ev.Summary = v.Value
		case "X-GOOGLE-CONFERENCE":
			ev.Link = v.Value
		}
	}

	m.Events = append(m.Events, &ev)

	return nil
}

func handleAttendee(ev *Event, a *ics.Attendee) {
	var att Attendee

	att.Email = a.Email()
	att.Status = a.ParticipationStatus()

	ev.Attendees = append(ev.Attendees, &att)
}
