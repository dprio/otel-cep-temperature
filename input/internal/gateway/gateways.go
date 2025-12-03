package gateway

import (
	"github.com/dprio/otel-cep-temperature/input/internal/gateway/weathergateway"
	"github.com/dprio/otel-cep-temperature/input/internal/infrastructure/httpclient"
)

type Gateways struct {
	WeatherGateway weathergateway.Gateway
}

func New(httpClients httpclient.HTTPClients) *Gateways {
	return &Gateways{
		WeatherGateway: weathergateway.New(httpClients.TempOrchestratorClient),
	}
}
