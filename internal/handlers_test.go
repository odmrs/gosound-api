package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Status))

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	// Verify Header
	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Header error, expected: %v\nGot %v", "application/json", resp.Header.Get("Content-Type"))
	}

	// Verify status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 but got %d", resp.StatusCode)
	}
	status := &statusCode{
		Code:    http.StatusOK,
		Status:  "Online",
		Message: "API is running smoothly",
	}

	// Verify JSON
	expected, _ := json.Marshal(status)
	b, err := io.ReadAll(resp.Body)

	if string(b) != string(expected) {
		t.Errorf("Expected: %v\nGot: %v", expected, err)
	}
}

func TestTts(t *testing.T) {
	jsonBody := `{"text": "Testing api"}`
	bodyRender := bytes.NewReader([]byte(jsonBody))
	server := httptest.NewServer(http.HandlerFunc(Tts))

	resp, err := http.Post(server.URL, "audio/mpeg", bodyRender)
	if err != nil {
		t.Error(err)
	}

	// Verify status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 status code, but got %d", resp.StatusCode)
	}

	// Verify audio length
	if resp.ContentLength == 0 {
		t.Error("Expected non-empty audio content")
	}

	// Verify Header
	if resp.Header.Get("Content-Type") != "audio/mpeg" {
		t.Errorf("Header error, expected: %v\nGot: %v", "audio/mpeg", resp.Header.Get("Content-Type"))
	}

	// Verify Content-Disposition"}
	if resp.Header.Get("Content-Disposition") != `attachment; filename="audio.mp3"` {
		t.Errorf("Content-Disposition expected: %v\nGot: %v", `attachment; filename="audio.mp3`, resp.Header.Get("Content-Disposition"))
	}
}
