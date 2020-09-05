package main

import (
	"log"

	"github.com/dorakueyon/go-filter/commands"
)

func main() {
	debug := true
	app, err := commands.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(debug)
	if err != nil {
		log.Fatal(err)
	}
	//execute(debug)
}
