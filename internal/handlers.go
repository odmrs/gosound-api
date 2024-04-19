package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var status string = "on"

type statusCode struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func StatusOn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if status != "on" {
		statusOff := &statusCode{
			Code:    http.StatusServiceUnavailable,
			Status:  "Offline",
			Message: "API is currently offline. Please try again later",
		}

		statusOffJson, _ := json.Marshal(statusOff)
		fmt.Fprintln(w, string(statusOffJson))
		return
	}

	statusOn := &statusCode{
		Code:    http.StatusOK,
		Status:  "Online",
		Message: "API is running smoothly",
	}

	statusOffJson, _ := json.Marshal(statusOn)
	w.Write(statusOffJson)
}
