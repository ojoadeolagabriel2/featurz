package main

import (
	"featurz/api"
	"featurz/app"
)

func main() {
	application := app.App{}
	application.InitializeRouter()
	api.AddRoutes(application.Router)
}