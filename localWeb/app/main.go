package app

import (
	"github.com/kaiyingzhang/Gogolang/localWeb/app"
	// "./config
)

func main() {

	// config := config.GetConfig()

	app := &app.App{}
	app.Initialize()
	app.Run(":8080")

}
