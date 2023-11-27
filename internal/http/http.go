package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

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

type response struct {
	Code     int `json:"code"`
	Iso639_3 int `json:"iso_693_3"`
}

func (c *myHttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "no text provded", http.StatusBadRequest)
		return
	}

	language, ok := c.detector.DetectLanguageOf(text)
	if !ok {
		http.Error(w, "cannot determine language", http.StatusBadRequest)
		return
	}

	resp := &response{
		Code:     int(language),
		Iso639_3: int(language.IsoCode639_3()),
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
