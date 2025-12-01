package gettemperaturebyzipcode

import (
	"context"
	"fmt"

	"github.com/dprio/otel-cep-temperature/internal/domain/address"
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

func (u *useCase) Execute(ctx context.Context, zipCode string) (*Output, error) {
	fmt.Printf("Executing GetTemperatureByZipCodeUseCase. [zipCode: %s]\n", zipCode)
	zc, err := address.NewZipCode(zipCode)
	if err != nil {
		return nil, err
	}

	addr, err := u.viaCEPGateway.GetAddressByZipCode(ctx, zc)
	if err != nil {
		return nil, err
	}

	wt, err := u.weatherAPIGateway.GetWeatherByCity(ctx, addr.City)
	if err != nil {
		return nil, err
	}

	return &Output{
		Address: *addr,
		Weather: *wt,
	}, nil
}
