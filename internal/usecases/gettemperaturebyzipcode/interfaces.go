package gettemperaturebyzipcode

import (
	"context"

	"github.com/dprio/cep-temperature/internal/domain/address"
	"github.com/dprio/cep-temperature/internal/domain/weather"
)

type (
	ViaCEPGateway interface {
		GetAddressByZipCode(ctx context.Context, zipCode address.ZipCode) (*address.Address, error)
	}

	WeatherAPIGateway interface {
		GetWeatherByCity(ctx context.Context, city string) (*weather.Weather, error)
	}

	UseCase interface {
		Execute(ctx context.Context, zipCode string) (*weather.Weather, error)
	}
)
