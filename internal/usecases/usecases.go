package usecases

import (
	"github.com/dprio/cep-temperature/internal/gateway"
	"github.com/dprio/cep-temperature/internal/usecases/gettemperaturebyzipcode"
)

type UseCases struct {
	GetTemperatureByZipCodeUseCase gettemperaturebyzipcode.UseCase
}

func New(gateways *gateway.Gateways) *UseCases {
	return &UseCases{
		GetTemperatureByZipCodeUseCase: gettemperaturebyzipcode.New(gateways.AddressGateway, gateways.WeatherGateway),
	}
}
