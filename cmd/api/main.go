package main

import (
	handlers "github.com/odmrs/gosound-api/internal"
	"log"
	"net/http"
	"os"
	"time"
)

const status string = "on"

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/gosoundapi", handlers.StatusOn)
	mux.HandleFunc("POST /v1/gosoundapi/tts", handlers.Tts)
  mux.HandleFunc("POST /v1/gosoundapi/stt", handlers.Stt)
  
	s := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting server on %s", s.Addr)
	err := http.ListenAndServe(":4000", mux)
	logger.Fatal(err)
}
