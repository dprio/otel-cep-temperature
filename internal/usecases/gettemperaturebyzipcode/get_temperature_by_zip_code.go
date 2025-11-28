package gettemperaturebyzipcode

import (
	"context"
	"fmt"

	"github.com/dprio/cep-temperature/internal/domain/address"
	"github.com/dprio/cep-temperature/internal/domain/weather"
)

type useCase struct {
	viaCEPGateway     ViaCEPGateway
	weatherAPIGateway WeatherAPIGateway
}

func New(viaCEPGateway ViaCEPGateway, weatWeatherAPIGateway WeatherAPIGateway) UseCase {
	return &useCase{
		viaCEPGateway:     viaCEPGateway,
		weatherAPIGateway: weatWeatherAPIGateway,
	}
}

func (u *useCase) Execute(ctx context.Context, zipCode string) (*weather.Weather, error) {
	fmt.Printf("Executing GetTemperatureByZipCodeUseCase. [zipCode: %s]\n", zipCode)
	zc, err := address.NewZipCode(zipCode)
	if err != nil {
		return nil, err
	}

	addr, err := u.viaCEPGateway.GetAddressByZipCode(ctx, zc)
	if err != nil {
		return nil, err
	}

	return u.weatherAPIGateway.GetWeatherByCity(ctx, addr.City)
}
