package main

import (
	"log"

	"metrics/internal/app"
)

func main() {
	if err := app.RunAuthApp(); err != nil {
		log.Fatal(err)
	}
}
