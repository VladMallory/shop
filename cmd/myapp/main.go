package main

import (
	"authTest/internal/app"
	"log"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
