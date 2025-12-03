package gettemperaturebyzipcode

import (
	"context"

	"github.com/dprio/otel-cep-temperature/input/internal/domain/weather"
)

type (
	WeatherGateway interface {
		GetWeatherByCep(ctx context.Context, cep string) (*weather.Weather, error)
	}

	UseCase interface {
		Execute(ctx context.Context, zipCode string) (*Output, error)
	}
)
