package gettemperaturebyzipcode

import (
	"github.com/dprio/otel-cep-temperature/input/internal/domain/weather"
)

type Output struct {
	Weather weather.Weather
}
