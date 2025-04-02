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
		{"re", lingua.Unknown},
		{"я кто", lingua.Russian},
		{"я хто", lingua.Ukrainian},
		{"привет", lingua.Russian},
		{"Привет, как твои дела?", lingua.Russian},
		{"漢語漢語", lingua.Chinese},
		{"Olá, como estás?", lingua.Spanish},
		{"Привіт, як твої справи?", lingua.Ukrainian},
		{"やあ、元気かい？", lingua.Japanese},
		{"Salut, comment ça va?", lingua.French},
	}

	for _, c := range cases {
		lang, _ := instance.DetectLanguageOf(c.text)

		if lang != c.want {
			t.Errorf(`"%s" got %v, want %v`, c.text, lang, c.want)
		}
	}
}
