package download

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func DownloadAudio(r *http.Request) string {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Panicf("Error parsing multipart form: %v", err)
	}

	file, _, err := r.FormFile("audio")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	filePath := filepath.Join(".", "pkg", "download", "audio_placeholder")
	err = os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		log.Panicf("Error creating audio_placeholder folder")
	}
	filePath += time.Now().Format("20060102150405.999999999")

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
