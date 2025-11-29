package gateway

import (
	"github.com/dprio/cep-temperature/internal/gateway/addressgateway"
	weathergateway "github.com/dprio/cep-temperature/internal/gateway/weatherGateway"
	"github.com/dprio/cep-temperature/internal/infrastructure/httpclient"
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
