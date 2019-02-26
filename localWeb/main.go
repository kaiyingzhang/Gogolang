package main

import (
	"github.com/kaiyingzhang/Gogolang/localWeb/app"
	// "./config
)

func main() {

	app := &app.App{}
	app.Initialize()
	app.Run(":8080")

}
