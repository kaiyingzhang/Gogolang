package main

import (
	"github.com/kaiyingzhang/Gogolang/tree/master/Go_HTTP/localWeb/app"
	// "./config"
)

func main() {

	// config := config.GetConfig()

	app := &app.App{}
	app.Initialize()
	app.Run(":8080")

}
