package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func apiHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ONLINE")
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/gosoundapi", apiHealth)

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
