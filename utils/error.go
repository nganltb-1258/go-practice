package utils

import (
    "encoding/json"
    "net/http"
    "log"
)

type ErrorResponse struct {
    Error string `json:"error"`
    Code  int    `json:"code"`
}

func WriteError(w http.ResponseWriter, message string, statusCode int) {
    log.Printf("Error: %s, Status Code: %d", message, statusCode)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)

    json.NewEncoder(w).Encode(ErrorResponse{
        Error: message,
        Code:  statusCode,
    })
}
