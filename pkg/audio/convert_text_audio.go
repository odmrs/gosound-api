package audio

import (
	"log"

	"github.com/hegedustibor/htgo-tts"

	"github.com/odmrs/gosound-api/pkg/file"
)

func ConvertTextToAudio(text string) (string, error) {
	// Convert audio to text
	speech := htgotts.Speech{Folder: "./pkg/audio/text_audio_placeholder", Language: "pt"}
	if err := speech.Speak(text); err != nil {
		log.Panicf("Error to try speak the text, error: %v", err)
	}

	// Return the file
	return getFile.GetFile(speech.Folder)
}
