package main

import (
	"restFul-ibnu/app"
	"restFul-ibnu/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
