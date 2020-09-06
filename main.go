package main

import (
	"log"

	"github.com/dorakueyon/go-filter/commands"
)

func main() {
	debug := true
	createOutputFile := true
	app, err := commands.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(debug, createOutputFile)
	if err != nil {
		log.Fatal(err)
	}
}
