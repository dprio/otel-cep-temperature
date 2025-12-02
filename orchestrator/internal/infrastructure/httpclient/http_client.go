package httpclient

import (
	"net/http"
	"time"

	"github.com/dprio/otel-cep-temperature/orchestrator/internal/infrastructure/httpclient/viacephttpclient"
	"github.com/dprio/otel-cep-temperature/orchestrator/internal/infrastructure/httpclient/weatherapihttpclient"
)

type HTTPClients struct {
	WeatherAPIHTTPClient weatherapihttpclient.Client
	ViaCEPHTTTPClient    viacephttpclient.Client
}

func New() *HTTPClients {
	return &HTTPClients{
		WeatherAPIHTTPClient: weatherapihttpclient.New(&http.Client{Timeout: time.Second * 10}),
		ViaCEPHTTTPClient:    viacephttpclient.New(&http.Client{Timeout: time.Second * 10}),
	}

}
