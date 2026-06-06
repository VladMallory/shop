package main

import (
	"authTest/internal/app"
	"log"
)

func main() {
	application, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	application.Run()
}
