package handlers

import (
	"encoding/json"
	"featurz/data"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

func GetAllFeatures(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	log.Printf("processing GetAllFeatures\n")
	writer.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(data.SampleData)
	_, _ = writer.Write(payload)
}

func GetFeatureById(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(request)
	key := vars["id"]

	for _, feature := range data.SampleData {
		if feature.Id == key {
			payload, err := json.Marshal(feature)
			if err != nil {
				fmt.Println(err.Error())
			}
			writer.Header().Set("Content-Type", "application/json")
			_, _ = writer.Write(payload)
			return
		}
	}

	emptyResponse, _ := json.Marshal(make(map[string]string))
	_, err := writer.Write(emptyResponse)
	if err != nil {
		log.Println(err.Error())
	}
}

func GetFeatureLikeName(writer http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(request)
	key := vars["name"]

	var selectedFeatures []data.Feature
	for _, feature := range data.SampleData {
		if strings.Contains(feature.Name, key) {
			selectedFeatures = append(selectedFeatures, feature)
		}
	}

	if len(selectedFeatures) != 0 {
		writer.Header().Set("Content-Type", "application/json")
		payload, _ := json.Marshal(selectedFeatures)
		_, _ = writer.Write(payload)
		return
	}

	emptyResponse, _ := json.Marshal(make(map[string]string))
	_, _ = writer.Write(emptyResponse)
}

func CreateFeature(writer http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	writer.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(data.TestFeatures{})
	_, _ = writer.Write(payload)
}
