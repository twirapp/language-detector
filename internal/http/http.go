package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strings"

	iso6391 "github.com/emvi/iso-639-1"
	"github.com/twirapp/language-detector/internal/predictor"
)

type myHttp struct {
	detector *predictor.Predictor
}

func New(ctx context.Context, port string, predictor *predictor.Predictor) {
	server := &http.Server{Addr: "0.0.0.0:" + port, Handler: &myHttp{
		detector: predictor,
	}}
	go func() {
		log.Printf("Starting listening on %s port\n", port)
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				if err := server.Shutdown(ctx); err != nil {
					panic(err)
				}

				log.Printf("Shutting down\n", port)
				break
			}
		}
	}()
}

type responseLang struct {
	Iso639_1   string `json:"iso_693_1"`
	Name       string `json:"name"`
	NativeName string `json:"native_name"`
}

func (c *myHttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/languages" {
		c.allLanguagesController(w, r)
		return
	}

	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "no text provded", http.StatusBadRequest)
		return
	}

	languages, err := c.detector.Predict(text)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	resp := make([]responseLang, len(languages))

	for i, lang := range languages {
		parsedLang := iso6391.FromCode(lang.Label)

		resp[i] = responseLang{
			Name:       parsedLang.Name,
			Iso639_1:   parsedLang.Code,
			NativeName: parsedLang.NativeName,
		}
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func (c *myHttp) allLanguagesController(w http.ResponseWriter, r *http.Request) {
	allLangs := iso6391.Languages
	resp := make([]responseLang, 0, len(allLangs))

	for _, lang := range allLangs {
		resp = append(
			resp,
			responseLang{
				Name:       lang.Name,
				Iso639_1:   lang.Code,
				NativeName: lang.NativeName,
			},
		)
	}

	slices.SortFunc(resp, func(a, b responseLang) int {
		return strings.Compare(a.Name, b.Name)
	})

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
