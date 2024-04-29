package handlers

import (
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
