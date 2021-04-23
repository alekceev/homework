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

	var port string
	flag.StringVar(&port, "port", "8081", "Server port")
	flag.Parse()

	app := app.NewApp()
	app.DB = &database.DB{}
	if err := app.DB.Open(); err != nil {
		log.Fatalf("Db error: %v", err)
		os.Exit(1)
	}
	defer app.DB.Close()
	app.InitRouters()

	http.HandleFunc("/", app.Router.ServeHTTP)

	fmt.Println("Server is listening http://localhost:" + port + "...")
	http.ListenAndServe(":"+port, nil)
}
