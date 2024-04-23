package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var status string = "on"

type statusCode struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type textToSpeech struct {
	Text string `json:"text"`
}

// Show the status of api
func StatusOn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if status != "on" {
		statusOff := &statusCode{
			Code:    http.StatusServiceUnavailable,
			Status:  "Offline",
			Message: "API is currently offline. Please try again later",
		}

		statusOffJson, _ := json.Marshal(statusOff)
		fmt.Fprintln(w, string(statusOffJson))
		return
	}

	statusOn := &statusCode{
		Code:    http.StatusOK,
		Status:  "Online",
		Message: "API is running smoothly",
	}

	statusOffJson, _ := json.Marshal(statusOn)
	w.Write(statusOffJson)
}

// Handler -> Text to Speech
func Tts(w http.ResponseWriter, r *http.Request) {
	var textDecoded textToSpeech
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Content-Disposition", "attachment; filename=\"audio.mp3\"")

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&textDecoded)
	if err != nil {
		log.Panic(err)
	}

	audioPath, err := convertTextToAudio(textDecoded.Text)
	if err != nil {
		log.Panic(err)
	}

	audioData, err := os.ReadFile(audioPath)
	if err != nil {
		http.Error(w, "Error to read audio file", http.StatusInternalServerError)
	}

	w.Write(audioData)
	log.Println("Audio sended")

	if err := os.Remove(audioPath); err != nil {
		log.Printf("Failed to remove audio file %v", err)
	}
}

func Stt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Content-Disposition", "attachment; filename=\"speech.mp3\"")

	downloadAudio(r)
}
