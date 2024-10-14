package main

import (
	"log"

	"metrics/internal/app"
)

func main() {
	if err := app.RunGeoApp(); err != nil {
		log.Fatal(err)
	}

}
