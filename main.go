package main

import (
	"clean-architecture/config"
	"clean-architecture/route"
)

func main() {
	config.SetEnv()
	app := route.NewService()
	if err := app.Start(); err != nil {
		panic(err)
	}
}
