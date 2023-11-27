package lingua

import (
	"testing"

	"github.com/pemistahl/lingua-go"
)

func TestNew(t *testing.T) {
	instance := New("development")

	lang, ok := instance.DetectLanguageOf("Привет как твои дела?")

	if !ok {
		t.Error("cannot determine language")
	}

	if lang != lingua.Russian {
		t.Error("language is not russian")
	}
}
