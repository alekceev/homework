package main

import (
	"log"
	"os"

	"homework/pkg/app"
	"homework/pkg/database"

	"github.com/namsral/flag"
)

func main() {

	var host, port, dbHost string

	flag.StringVar(&host, "host", "localhost", "Server host")
	flag.StringVar(&port, "port", "8081", "Server port")
	flag.StringVar(&dbHost, "dbHost", "file:/tmp/catalog.db?_mutex=full&_cslike=false", "Server dbHost")
	flag.Parse()

	db, err := database.Connect(dbHost)
	if err != nil {
		log.Fatalf("Db error: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	app := app.NewApp()
	app.DB = db

	if err := app.Start(host, port); err != nil {
		panic(err)

		// gracefull shutown
	}

}
