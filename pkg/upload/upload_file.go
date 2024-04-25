package uploadHandler

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadFile(path string) ([]byte, error) {
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
