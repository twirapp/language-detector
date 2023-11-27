package lingua

import (
	l "github.com/pemistahl/lingua-go"
)

type Lingua struct {
	l.LanguageDetector
}

func New(appEnv string) *Lingua {
	var detector l.LanguageDetector
	if appEnv == "development" {
		detector = l.NewLanguageDetectorBuilder().
			FromAllLanguages().
			WithLowAccuracyMode().
			Build()
	} else {
		detector = l.NewLanguageDetectorBuilder().
			FromAllLanguages().
			Build()
	}

	return &Lingua{
		LanguageDetector: detector,
	}
}
