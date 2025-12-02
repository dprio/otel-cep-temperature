package main

import "github.com/dprio/otel-cep-temperature/orchestrator/cmd/ceptemporchestration/app"

func main() {
	app := app.New()

	if err := app.Start(); err != nil {
		panic(err)
	}
}
