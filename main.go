package main

import (
	"homework/pkg/app"
)

func main() {
	app := app.NewApp()
	defer app.Close()

	if err := app.Start(); err != nil {
		panic(err)

		// gracefull shutown
	}
}
