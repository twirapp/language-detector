package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	l "github.com/pemistahl/lingua-go"
	"github.com/twirapp/language-detector/internal/lingua"
)

type myHttp struct {
	detector *lingua.Lingua
}

func New(ctx context.Context, port string, lingua *lingua.Lingua) {
	server := &http.Server{Addr: "0.0.0.0:" + port, Handler: &myHttp{
		detector: lingua,
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
	Code     int    `json:"code"`
	Iso639_3 int    `json:"iso_693_3"`
	Name     string `json:"name"`
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

	languages := c.detector.DetectMultipleLanguagesOf(text)

	resp := make([]responseLang, len(languages))

	for i, lang := range languages {
		resp[i] = responseLang{
			Code:     int(lang.Language()),
			Iso639_3: int(lang.Language().IsoCode639_3()),
			Name:     lang.Language().String(),
		}
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func (c *myHttp) allLanguagesController(w http.ResponseWriter, r *http.Request) {
	resp := make([]responseLang, len(l.AllLanguages()))

	for i, lang := range l.AllLanguages() {
		resp[i] = responseLang{
			Code:     int(lang),
			Iso639_3: int(lang.IsoCode639_3()),
			Name:     lang.String(),
		}
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
