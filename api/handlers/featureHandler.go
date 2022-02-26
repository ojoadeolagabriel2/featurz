package handlers

import (
	"encoding/json"
	"featurz/data"
	"net/http"
)

func GetAllFeatures(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(data.TestFeatures{})
	_, _ = writer.Write(payload)
}

func GetFeatureById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(data.TestFeatures{})
	_, _ = writer.Write(payload)
}
