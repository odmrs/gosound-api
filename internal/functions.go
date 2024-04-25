package handlers

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hegedustibor/htgo-tts"
)

func convertTextToAudio(text string) (string, error) {
	// Convert audio to text
	speech := htgotts.Speech{Folder: "./internal/audio", Language: "pt"}
	if err := speech.Speak(text); err != nil {
		log.Panicf("Error to try speak the text, error: %v", err)
	}

	// Return the file
	return getFile(speech.Folder)
}

// Get the last file inputed and return the file
func getFile(dir string) (string, error) {
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

// TODO Improve this code
func downloadAudio(r *http.Request) string {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Panicf("Error parsing multipart form: %v", err)
	}

	file, _, err := r.FormFile("audio")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	filePath := "./internal/speech/" + "audio"

	out, err := os.Create(filePath)
	if err != nil {
		log.Panic(err)
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Download file with success")
	return filePath
}

func uploadFile(path string) ([]byte, error) {
	var targetUrl string = "http://localhost:5000/transcribe"
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		log.Panic(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Panic(err)
	}

	writer.Close()

	// Send to api python
	request, err := http.NewRequest(http.MethodPost, targetUrl, body)
	if err != nil {
		log.Panic(err)
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())

	// Send request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Panic(err)
	}

	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Panic(err)
	}

	// Delete user audio
	log.Println(path)
	if err := os.Remove(path); err != nil {
		log.Printf("Failed to remove audio file %v", err)
	}
	return respBody, nil
}
