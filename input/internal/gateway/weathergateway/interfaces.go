package weathergateway

import (
	"context"

	"github.com/dprio/otel-cep-temperature/input/internal/domain/weather"
	"github.com/dprio/otel-cep-temperature/input/internal/infrastructure/httpclient/temporchestratorclient"
)

type (
	Client interface {
		GetCityWeatherInformation(ctx context.Context, city string) (*temporchestratorclient.Response, error)
	}

	Gateway interface {
		GetWeatherByCep(ctx context.Context, cep string) (*weather.Weather, error)
	}
)
