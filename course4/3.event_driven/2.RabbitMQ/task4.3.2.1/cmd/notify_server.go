package main

import (
	"log"

	"metrics/internal/app"
)

func main() {
	if err := app.RunNotifyApp(); err != nil {
		log.Fatal(err)
	}
}
