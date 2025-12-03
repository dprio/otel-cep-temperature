package handlers

import (
	"github.com/dprio/otel-cep-temperature/input/internal/infrastructure/web/handlers/weatherhandler"
	"github.com/dprio/otel-cep-temperature/input/internal/usecases"
)

type Handlers struct {
	WeatherHandler weatherhandler.WeatherHandler
}

func New(useCases *usecases.UseCases) *Handlers {
	return &Handlers{
		WeatherHandler: *weatherhandler.New(useCases.GetTemperatureByZipCodeUseCase),
	}
}
