package api

import (
	"encoding/json"
	"featurz/api/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ojoadeolagabriel2/autogate-go-core/util"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

const (
	EnvApiVersionPrefixKey = "ENV_API_VERSION_PREFIX_KEY"
	EnvAppPort             = "ENV_APP_PORT"
)

var EnvApiPrefixKey = util.GetConfig(EnvApiVersionPrefixKey, "/v1")
var DefaultPort = util.GetIntConfig(EnvAppPort, 12345)

func AddRoutes(router *mux.Router) {
	fmt.Println("\nFeature Service Rest API - v1.0\n===============================")

	// defaults
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		log.Println("/health endpoint called")
	}).Methods("GET")

	// user handler mapping
	router.
		Methods(http.MethodGet).
		Path(EnvApiPrefixKey+"/features").
		Queries("name", "{name:[a-zA-Z]+}").
		Handler(negroni.New(negroni.HandlerFunc(handlers.GetFeatureLikeName)))
	log.Println("handler [GetFeatureLikeName] registered")

	router.
		Methods(http.MethodGet).
		Path(EnvApiPrefixKey+"/features").
		Queries("tag", "{name:[a-zA-Z]+}").
		Handler(negroni.New(negroni.HandlerFunc(handlers.GetFeatureByTag)))
	log.Println("handler [GetFeatureByTag] registered")

	router.
		Methods(http.MethodPut).
		Path(EnvApiPrefixKey+"/features").
		Handler(negroni.New(negroni.HandlerFunc(handlers.UpdateFeature)))
	log.Println("handler [UpdateFeature] registered")

	router.
		Methods(http.MethodGet).
		Path(EnvApiPrefixKey + "/features/{id}").
		Handler(negroni.New(negroni.HandlerFunc(handlers.GetFeatureById)))
	log.Println("handler [GetFeatureById] registered")

	router.
		Methods(http.MethodGet).
		Path(EnvApiPrefixKey + "/features").
		Handler(negroni.New(negroni.HandlerFunc(handlers.GetAllFeatures)))
	log.Println("handler [GetAllFeatures] registered")

	router.
		Methods(http.MethodPost).
		Path(EnvApiPrefixKey + "/features").
		Handler(negroni.New(negroni.HandlerFunc(handlers.CreateFeature)))
	log.Println("handler [CreateFeature] registered")

	err := http.ListenAndServe(fmt.Sprintf(":%d", util.GetIntConfig(EnvAppPort, DefaultPort)), router)
	if err != nil {
		log.Fatalln("test failed: " + err.Error())
	}
}
