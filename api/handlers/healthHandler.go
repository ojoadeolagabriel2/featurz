package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetHealthStatus(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	log.Printf("processing GetHealthStatus\n")
	writer.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(map[string]string{"status": "ok", "db": "ok"})
	_, _ = writer.Write(payload)
}