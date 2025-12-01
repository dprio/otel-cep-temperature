package gettemperaturebyzipcode

import (
	"github.com/dprio/otel-cep-temperature/internal/domain/address"
	"github.com/dprio/otel-cep-temperature/internal/domain/weather"
)

type Output struct {
	Address address.Address
	Weather weather.Weather
}
