package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/odmrs/gosound-api/pkg/audio"
	"github.com/odmrs/gosound-api/pkg/download"
	"github.com/odmrs/gosound-api/pkg/upload"
)

type statusCode struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type textToSpeech struct {
	Text string `json:"text"`
}

// Show the status of api
func Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	status := &statusCode{
		Code:    http.StatusOK,
		Status:  "Online",
		Message: "API is running smoothly",
	}

	statusJson, _ := json.Marshal(status)
	if _, err := w.Write(statusJson); err != nil {
		log.Panicf("Error to parssing statusJson to json, error: %v", err)
	}
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

	audioPath, err := audio.ConvertTextToAudio(textDecoded.Text)
	if err != nil {
		log.Panic(err)
	}

	audioData, err := os.ReadFile(audioPath)
	if err != nil {
		http.Error(w, "Error to read audio file", http.StatusInternalServerError)
	}

	if _, err := w.Write(audioData); err != nil {
		log.Panicf("Error to try send audio to user, error: %v", err)
	}
	log.Println("Audio sended")

	if err := os.Remove(audioPath); err != nil {
		log.Printf("Failed to remove audio file %v", err)
	}
}

func Stt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Content-Disposition", "attachment; filename=\"speech.mp3\"")

	path := download.DownloadAudio(r)
	jsonResponse, err := uploadHandler.UploadFile(path)
	if err != nil {
		log.Panic(err)
	}
	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(jsonResponse); err != nil {
		log.Panicf("Error to try return json response to user, error: %v", err)
	}
}
