package main

import "github.com/dprio/otel-cep-temperature/input/cmd/app"

func main() {
	app := app.New()

	if err := app.Start(); err != nil {
		panic(err)
	}
}
