package main

import (
	"log"
	"net/http"
	"time"

	handlers "github.com/odmrs/gosound-api/internal"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/gosoundapi", handlers.Status)
	mux.HandleFunc("POST /v1/gosoundapi/tts", handlers.Tts)
	mux.HandleFunc("POST /v1/gosoundapi/stt", handlers.Stt)

	s := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Starting server on %s", s.Addr)
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
