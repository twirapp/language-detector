package lingua

import (
	"testing"

	"github.com/pemistahl/lingua-go"
)

func TestNew(t *testing.T) {
	instance := New("development")

	cases := []struct {
		text string
		want lingua.Language
	}{
		{"re how are you", lingua.English},
		{"я кто", lingua.Russian},
		{"привет", lingua.Russian},
		{"Привет, как твои дела?", lingua.Russian},
		{"漢語漢語", lingua.Chinese},
	}

	for _, c := range cases {
		lang, ok := instance.DetectLanguageOf(c.text)

		if !ok {
			t.Errorf(`"%s" got %v, want %v`, c.text, ok, true)
			continue
		}

		if lang != c.want {
			t.Errorf(`"%s" got %v, want %v`, c.text, lang, c.want)
		}
	}
}
