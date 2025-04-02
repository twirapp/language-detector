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
		WithMinimumRelativeDistance(0.01).
		Build()

	return &Lingua{
		LanguageDetector: detector,
	}
}
