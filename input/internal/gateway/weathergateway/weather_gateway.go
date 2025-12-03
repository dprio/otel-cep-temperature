package weathergateway

import (
	"context"

	"github.com/dprio/otel-cep-temperature/input/internal/domain/weather"
)

type weatherGateway struct {
	client Client
}

func New(client Client) Gateway {
	return &weatherGateway{
		client: client,
	}
}

func (g *weatherGateway) GetWeatherByCep(ctx context.Context, cep string) (*weather.Weather, error) {
	resp, err := g.client.GetCityWeatherInformation(ctx, cep)
	if err != nil {
		return nil, err
	}

	return weather.New(resp.City, resp.CelsiusTemp, resp.FahrenheitTemp, resp.KelvinTemperature), nil
}
