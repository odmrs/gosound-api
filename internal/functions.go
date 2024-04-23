package handlers

import (
	"os"
	"path/filepath"

	"github.com/hegedustibor/htgo-tts"
)

func convertTextToAudio(text string) (string, error){
  // Convert audio to text
  speech := htgotts.Speech{Folder: "./internal/audio", Language: "pt"}
	speech.Speak(text)

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
