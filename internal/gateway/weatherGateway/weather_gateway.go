package weathergateway

import (
	"context"

	"github.com/dprio/cep-temperature/internal/domain/weather"
)

type weatherGateway struct {
	client Client
}

func New(client Client) Gateway {
	return &weatherGateway{
		client: client,
	}
}

func (g *weatherGateway) GetWeatherByCity(ctx context.Context, city string) (*weather.Weather, error) {
	resp, err := g.client.GetCityWeatherInformation(ctx, city)
	if err != nil {
		return nil, err
	}

	return &weather.Weather{
		Temperature: *weather.NewTemparature(resp.Current.TempC),
	}, nil
}
