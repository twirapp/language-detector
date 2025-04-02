package predictor

import (
	"os"
	"path/filepath"

	"github.com/nano-interactive/go-fasttext"
)

func New() (*Predictor, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	ff, err := fasttext.Open(filepath.Join(cwd, "lid.176.bin"))
	if err != nil {
		return nil, err
	}

	return &Predictor{
		ff: ff,
	}, nil
}

type Predictor struct {
	ff fasttext.Model
}

type Prediction struct {
	Label       string  `json:"label"`
	Probability float32 `json:"probability"`
}

func (p *Predictor) Predict(text string) ([]Prediction, error) {
	predictions, err := p.ff.Predict(text, 1, 0)
	if err != nil {
		return nil, err
	}

	pr := make([]Prediction, len(predictions))
	for i := range predictions {
		pr[i] = Prediction{
			Label:       predictions[i].Label,
			Probability: predictions[i].Probability,
		}
	}

	return pr, nil
}
