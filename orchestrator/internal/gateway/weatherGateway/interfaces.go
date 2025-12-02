package weathergateway

import (
	"context"

	"github.com/dprio/otel-cep-temperature/orchestrator/internal/domain/weather"
	"github.com/dprio/otel-cep-temperature/orchestrator/internal/infrastructure/httpclient/weatherapihttpclient"
)

type (
	Client interface {
		GetCityWeatherInformation(ctx context.Context, city string) (*weatherapihttpclient.Response, error)
	}

	Gateway interface {
		GetWeatherByCity(ctx context.Context, city string) (*weather.Weather, error)
	}
)
