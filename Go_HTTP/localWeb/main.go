package main

import (
	"./app"
	// "./config"
)

func main() {

	// config := config.GetConfig()

	app := &app.App{}
	app.Initialize()
	app.Run(":8080")

}
