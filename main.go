package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"homework/pkg/app"
	"homework/pkg/database"

	"github.com/namsral/flag"
)

func main() {

	// todo: move to flags
	var dbHost = "file:/tmp/catalog.db?_mutex=full&_cslike=false"

	var port string
	flag.StringVar(&port, "port", "8081", "Server port")
	flag.Parse()

	db, err := database.Connect(dbHost)
	if err != nil {
		log.Fatalf("Db error: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	app := app.NewApp()
	app.DB = db

	app.InitRouters()

	http.HandleFunc("/", app.Router.ServeHTTP)

	fmt.Println("Server is listening http://localhost:" + port + "...")
	http.ListenAndServe(":"+port, nil)
}
