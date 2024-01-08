package email

import (
	"testing"

	"github.com/mnako/letters"
)

func TestGetFileName(t *testing.T) {
	contentType := letters.ContentTypeHeader{
		Params: map[string]string{
			"filename": "test.txt",
		},
	}

	disposition := letters.ContentDispositionHeader{
		Params: map[string]string{
			"name": "test2.txt",
		},
	}

	filename := getFileName(contentType, disposition)
	if filename != "test.txt" {
		t.Errorf("Expected filename to be 'test.txt', got '%s'", filename)
	}

	contentType.Params = map[string]string{}
	filename = getFileName(contentType, disposition)
	if filename != "test2.txt" {
		t.Errorf("Expected filename to be 'test2.txt', got '%s'", filename)
	}

	disposition.Params = map[string]string{}
	filename = getFileName(contentType, disposition)
	if filename != "" {
		t.Errorf("Expected filename to be '', got '%s'", filename)
	}
}
