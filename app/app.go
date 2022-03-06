package app

import (
	"database/sql"
	"featurz/data"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router  *mux.Router
	DB      *sql.DB
	Storage *data.FeatureStorage
}

type DbConfigOpt struct {
	Username     string
	Password     string
	Host         string
	Port         int
	DatabaseName string
}

func (app *App) InitializeRouter() {
	router := mux.NewRouter().StrictSlash(true)
	app.Router = router
}

func (app *App) InitializeDb(opt DbConfigOpt) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		opt.Host, opt.Port, opt.Username, opt.Password, opt.DatabaseName)

	// open connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// ping
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Connected to database @ localhost:%d/%s\n", opt.Port, opt.DatabaseName)
}
