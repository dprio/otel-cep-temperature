package viacephttpclient

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dprio/otel-cep-temperature/internal/infrastructure/httpclient/client"
)

var (
	ErrMakingRequestViaCEPAPI = errors.New("error making request to ViaCEP API")
	ErrViaCEPAPI              = errors.New("ViaCEPAPI returned an error")
	ErrCreatingRequest        = errors.New("error creating request to ViaCEPAPI")
)

type Client interface {
	GetAddress(ctx context.Context, cep string) (*Response, error)
}

type viacepClient struct {
	httpClient client.HttpClient
	basePath   string
}

func New(httpClient client.HttpClient) Client {
	basePath := "https://viacep.com.br/ws/"

	return &viacepClient{
		httpClient: httpClient,
		basePath:   basePath,
	}
}

func (c *viacepClient) GetAddress(ctx context.Context, cep string) (*Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.basePath+cep+"/json/", nil)
	if err != nil {
		return nil, ErrCreatingRequest
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, ErrMakingRequestViaCEPAPI
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrViaCEPAPI
	}

	var response Response
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
