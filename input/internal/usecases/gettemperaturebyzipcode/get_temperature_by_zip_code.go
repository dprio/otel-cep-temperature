package gettemperaturebyzipcode

import (
	"context"
	"fmt"
)

type useCase struct {
	weatherAPIGateway WeatherGateway
}

func New(weatherGateway WeatherGateway) UseCase {
	return &useCase{
		weatherAPIGateway: weatherGateway,
	}
}

func (u *useCase) Execute(ctx context.Context, zipCode string) (*Output, error) {
	fmt.Printf("Executing GetTemperatureByZipCodeUseCase. [zipCode: %s]\n", zipCode)

	wt, err := u.weatherAPIGateway.GetWeatherByCep(ctx, zipCode)
	if err != nil {
		return nil, err
	}

	return &Output{
		Weather: *wt,
	}, nil
}
