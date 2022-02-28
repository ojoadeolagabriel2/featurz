package main

import (
	"featurz/api"
	"featurz/app"
	"github.com/ojoadeolagabriel2/autogate-go-core/util"
)

func main() {
	application := app.App{}
	application.InitializeRouter()
	application.InitializeDb(
		util.GetConfig("ENV_POSTGRES_USERNAME", "postgres"),
		util.GetConfig("ENV_POSTGRES_PASSWORD", "postgres"),
	)

	// init api routes
	api.AddRoutes(application.Router)
}
