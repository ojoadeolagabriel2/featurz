package main

import (
	"encoding/json"
	"featurz/api/handlers"
	"featurz/app"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ojoadeolagabriel2/autogate-go-core/util"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

const (
	EnvAppPort = "ENV_APP_PORT"
)

var EnvApiPrefixKey = util.GetConfig("ENV_API_VERSION_PREFIX_KEY", "/v1")

func main() {
	application := app.App{}
	application.InitializeRouter()
	addRoutes(application.Router)
}

func addRoutes(router *mux.Router) {
	fmt.Println("Feature Service Rest API - v1.0")

	// defaults
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")

	// user handler mapping
	router.
		Methods(http.MethodGet).
		Path(EnvApiPrefixKey+"/features").
		Queries("name", "{name:[a-z,]+}").
		Handler(negroni.New(negroni.HandlerFunc(handlers.GetFeatureLikeName)))

	router.
		Methods(http.MethodGet).
		Path(EnvApiPrefixKey+"/features/{id}").
		Handler(negroni.New(negroni.HandlerFunc(handlers.GetFeatureById)))

	router.
		Methods(http.MethodGet).
		Path(EnvApiPrefixKey+"/features").
		Handler(negroni.New(negroni.HandlerFunc(handlers.GetAllFeatures)))

	router.
		Methods(http.MethodPost).
		Path(EnvApiPrefixKey+"/features").
		Handler(negroni.New(negroni.HandlerFunc(handlers.CreateFeature)))

	err := http.ListenAndServe(fmt.Sprintf(":%d", util.GetIntConfig(EnvAppPort, 12345)), router)
	if err != nil {
		log.Fatalln("test failed: " + err.Error())
	}
}