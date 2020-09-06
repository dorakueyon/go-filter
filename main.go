package main

import (
	"log"

	"github.com/dorakueyon/go-filter/commands"
)

func main() {
	app, err := commands.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
