package main

import (
	"featurz/api"
	"featurz/app"
	"featurz/values"
	"github.com/ojoadeolagabriel2/autogate-go-core/util"
)

func main() {
	initiate()
}

func initiate() {
	application := app.App{}
	application.InitializeRouter()

	//ini postgres driver
	application.InitializeDb(buildDatabaseOptions())

	// init api routes
	api.AddRoutes(application.Router)
}

func buildDatabaseOptions() app.DbConfigOpt {
	opt := app.DbConfigOpt{
		Username:     util.GetConfig(values.EnvPostgresUsername, "postgres"),
		Password:     util.GetConfig(values.EnvPostgresPassword, "postgres"),
		Host:         util.GetConfig(values.EnvPostgresHost, "127.0.0.1"),
		Port:         util.GetIntConfig(values.EnvPostgresUsername, 5438),
		DatabaseName: util.GetConfig(values.EnvPostgresUsername, "postgres"),
	}
	return opt
}
