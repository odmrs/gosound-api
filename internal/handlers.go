package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/odmrs/gosound-api/pkg/audio"
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
		err := fmt.Errorf(err.Error(), "Error to parssing statusJson to json")
		fmt.Println(err.Error())
	}
}

// Handler -> Text to Speech
func Tts(w http.ResponseWriter, r *http.Request) {
	var textDecoded textToSpeech
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Content-Disposition", "attachment; filename=\"audio.mp3\"")
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("no body in request")
		fmt.Println(err.Error())
		json.NewEncoder(w).Encode(audio.NewBadRequestError("the request body is empty; use the json tag 'text' with the value you want to hear"))
		return
	}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&textDecoded)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("error when trying to decode the Text sent by the user")
		fmt.Println(err.Error())
	}

	audioPath, err := audio.ConvertTextToAudio(textDecoded.Text)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("error to trying convert text to audio")
		fmt.Println(err.Error())
	}

	audioData, err := os.ReadFile(audioPath)
	if err != nil {
		http.Error(w, "Error to read audio file", http.StatusInternalServerError)
	}

	if _, err := w.Write(audioData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("error when trying to send audio to user")
		fmt.Println(err.Error())
	}
	log.Println("Audio sended")

	// Remove text -> audio
	if err := os.Remove(audioPath); err != nil {
		log.Printf("Failed to remove audio file %v", err)
	}
}

func Stt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Content-Disposition", "attachment; filename=\"speech.mp3\"")

	path := audio.DownloadAudio(r, w)
	jsonResponse, err := audio.UploadFile(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("error when trying send audio to api python")
		fmt.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(jsonResponse); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("error when trying to send jsonResponse to user")
		fmt.Println(err.Error())
	}
}
