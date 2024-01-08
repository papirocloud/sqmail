package email

import (
	"bytes"
	"encoding/base64"

	"github.com/mnako/letters"
)

func getFileName(contentType letters.ContentTypeHeader, disposition letters.ContentDispositionHeader) string {
	for k, v := range contentType.Params {
		if k == "filename" {
			return v
		}
	}

	for k := range disposition.Params {
		if k == "name" {
			return contentType.Params[k]
		}
	}

	return ""
}

func handleAttachments(m *Message, msg letters.Email) error {
	for _, a := range msg.AttachedFiles {
		if a.ContentType.ContentType == "application/ics" || a.ContentType.ContentType == "text/calendar" {
			if err := handleInvitation(m, bytes.NewReader(a.Data)); err != nil {
				return err
			}
			continue
		}

		m.HasAttachments = true

		data := bytes.NewBuffer(nil)
		_, err := base64.NewEncoder(base64.StdEncoding, data).Write(a.Data)
		if err != nil {
			return err
		}

		m.Attachments = append(m.Attachments, &Attachment{
			ContentType: a.ContentType.ContentType,
			Content:     data.Bytes(),
			Filename:    getFileName(a.ContentType, a.ContentDisposition),
		})
	}

	return nil
}

func handleEmbeds(m *Message, msg letters.Email) error {
	for _, e := range msg.InlineFiles {
		m.HasEmbeds = true

		data := bytes.NewBuffer(nil)
		_, err := base64.NewEncoder(base64.StdEncoding, data).Write(e.Data)
		if err != nil {
			return err
		}

		m.Embedded = append(m.Embedded, &EmbeddedFile{
			ContentType: e.ContentType.ContentType,
			ContentId:   e.ContentID,
			Content:     data.Bytes(),
			Filename:    getFileName(e.ContentType, e.ContentDisposition),
		})
	}

	return nil
}
