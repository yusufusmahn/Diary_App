package utils

import (
    "encoding/json"
    "net/http"
)

// respondJSON sends a structured JSON response with given status and payload
func RespondJSON(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

// respondError simplifies sending an error message as JSON
func RespondError(w http.ResponseWriter, status int, message string) {
    RespondJSON(w, status, map[string]string{"error": message})
}
