package audio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/hegedustibor/htgo-tts"
)

type RestErr struct {
	Message string `json:"message"`
	Err     string `json:"error,omitempty"`
	Code    int    `json:"code"`
}

// Bad Request Error

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

// Convert audio to text
func ConvertTextToAudio(text string) (string, error) {
	speech := htgotts.Speech{Folder: "./pkg/audio/placeholder", Language: "pt"}
	if err := speech.Speak(text); err != nil {
		err := fmt.Errorf("Error to try speak the text")
		fmt.Println(err.Error())
	}

	// Return the file
	return GetFile(speech.Folder)
}

func DownloadAudio(r *http.Request, w http.ResponseWriter) string {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("Error parsing multipart form")
		fmt.Println(err.Error())
	}

	file, _, err := r.FormFile("audio")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(NewBadRequestError("the request body is empty; use the 'audio' name with the value of your audio as formularie"))
	}
	defer file.Close()

	filePath := filepath.Join(".", "pkg", "audio", "placeholder")
	err = os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("Error creating audio_placeholder folder")
		fmt.Println(err.Error())
	}
	filePath += time.Now().Format("20060102150405.999999999")

	out, err := os.Create(filePath)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("error when trying create the file the file to download")
		fmt.Println(err.Error())
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("error when trying copy the file the file to download")
		fmt.Println(err.Error())
	}

	log.Println("Download file with success")
	return filePath
}

// Get the last file inputed and return the file
func GetFile(dir string) (string, error) {
	var file string

	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	entry := dirEntries[0]
	fileInfo, err := entry.Info()
	if err != nil {
		return "", err
	}
	file = filepath.Join(dir, fileInfo.Name())

	if file == "" {
		return "", os.ErrNotExist
	}

	return file, nil
}

func UploadFile(path string) ([]byte, error) {
	var targetUrl string = "http://api-python:5000/transcribe"
	file, err := os.Open(path)
	if err != nil {
		err := fmt.Errorf("error when try open the speech file")
		fmt.Println(err.Error())
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		err := fmt.Errorf("error when trying create the form file")
		fmt.Println(err.Error())
	}

	_, err = io.Copy(part, file)
	if err != nil {
		err := fmt.Errorf("error when try the copy file in after create the form file")
		fmt.Println(err.Error())
	}

	writer.Close()

	// Send to api python
	request, err := http.NewRequest(http.MethodPost, targetUrl, body)
	if err != nil {
		err := fmt.Errorf("error when trying send the file for the python api")
		fmt.Println(err.Error())
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())

	// Send request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		err := fmt.Errorf("error when trying send the request to python")
		fmt.Println(err.Error())
	}

	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		err := fmt.Errorf("error when trying read the body response")
		fmt.Println(err.Error())
	}

	// Delete user audio
	if err := os.Remove(path); err != nil {
		log.Printf("Failed to remove audio file %v", err)
	}
	return respBody, nil
}
