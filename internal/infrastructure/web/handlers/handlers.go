package handlers

import (
	"github.com/dprio/cep-temperature/internal/infrastructure/web/handlers/weatherhandler"
	"github.com/dprio/cep-temperature/internal/usecases"
)

type Handlers struct {
	WeatherHandler weatherhandler.WeatherHandler
}

func New(useCases *usecases.UseCases) *Handlers {
	return &Handlers{
		WeatherHandler: *weatherhandler.New(useCases.GetTemperatureByZipCodeUseCase),
	}
}
