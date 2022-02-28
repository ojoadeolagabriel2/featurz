package app

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) InitializeRouter() {
	router := mux.NewRouter().StrictSlash(true)
	app.Router = router
}

func (app *App) InitializeDb(username string, password string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, "postgres")

	var err error
	app.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("connected to postgres instance")
	}
}