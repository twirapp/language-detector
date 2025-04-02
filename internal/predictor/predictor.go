package predictor

import (
	"github.com/nano-interactive/go-fasttext"
)

func New(modelPath string) (*Predictor, error) {
	ff, err := fasttext.Open(modelPath)
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
