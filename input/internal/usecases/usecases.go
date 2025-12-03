package usecases

import (
	"github.com/dprio/otel-cep-temperature/input/internal/gateway"
	"github.com/dprio/otel-cep-temperature/input/internal/usecases/gettemperaturebyzipcode"
)

type UseCases struct {
	GetTemperatureByZipCodeUseCase gettemperaturebyzipcode.UseCase
}

func New(gateways *gateway.Gateways) *UseCases {
	return &UseCases{
		GetTemperatureByZipCodeUseCase: gettemperaturebyzipcode.New(gateways.WeatherGateway),
	}
}
