package main

import (
	"log"

	"metrics/internal/app"
)

func main() {
	if err := app.RunUserApp(); err != nil {
		log.Fatal(err)
	}
}
