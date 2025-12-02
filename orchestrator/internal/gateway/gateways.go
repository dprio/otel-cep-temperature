package gateway

import (
	"github.com/dprio/otel-cep-temperature/orchestrator/internal/gateway/addressgateway"
	weathergateway "github.com/dprio/otel-cep-temperature/orchestrator/internal/gateway/weatherGateway"
	"github.com/dprio/otel-cep-temperature/orchestrator/internal/infrastructure/httpclient"
)

type Gateways struct {
	AddressGateway addressgateway.Gateway
	WeatherGateway weathergateway.Gateway
}

func New(httpClients httpclient.HTTPClients) *Gateways {
	return &Gateways{
		AddressGateway: addressgateway.New(httpClients.ViaCEPHTTTPClient),
		WeatherGateway: weathergateway.New(httpClients.WeatherAPIHTTPClient),
	}
}
