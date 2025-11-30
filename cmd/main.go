package main

import "github.com/dprio/cep-temperature/cmd/app"

func main() {
	app := app.New()

	if err := app.Start(); err != nil {
		panic(err)
	}
}
