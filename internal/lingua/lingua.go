package lingua

import (
	l "github.com/pemistahl/lingua-go"
)

type Lingua struct {
	l.LanguageDetector
}

func New(appEnv string) *Lingua {
	detector := l.NewLanguageDetectorBuilder().
		FromAllLanguages().
		Build()

	return &Lingua{
		LanguageDetector: detector,
	}
}
