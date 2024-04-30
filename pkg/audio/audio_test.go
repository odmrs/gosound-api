package audio

import (
	"os"
	"testing"
)

// Test Convert audio to text and GetFile
func TestConvertTextToAudio(t *testing.T) {
	speech, err := ConvertTextToAudio("Testando")
	if err != nil {
		t.Error("Expected ConvertTextToAudio convert")
	}

	if speech == "" {
		t.Error("Expected return of ConvertTextToAudio is a string")
	}

	if err := os.RemoveAll("./pkg"); err != nil {
		t.Errorf("Error try remove trash files, error: %v", err)
	}
}
