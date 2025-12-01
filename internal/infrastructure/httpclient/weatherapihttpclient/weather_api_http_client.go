package weatherapihttpclient

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dprio/otel-cep-temperature/internal/infrastructure/httpclient/client"
)

var (
	ErrMakingRequestWeatherAPI = errors.New("error making request to Weather API")
	ErrWeatherAPI              = errors.New("Weather API returned an error")
	ErrCreatingRequest         = errors.New("error creating request to  Weather API")
)

type (
	Client interface {
		GetCityWeatherInformation(ctx context.Context, city string) (*Response, error)
	}

	weatherAPIClient struct {
		httpClient client.HttpClient
		basePath   string
	}
)

func New(httpClient client.HttpClient) Client {
	basePath := "https://api.weatherapi.com/v1"

	return &weatherAPIClient{
		httpClient: httpClient,
		basePath:   basePath,
	}
}

func (c *weatherAPIClient) GetCityWeatherInformation(ctx context.Context, city string) (*Response, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.basePath+"/current.json", nil)
	if err != nil {
		return nil, ErrCreatingRequest
	}

	q := req.URL.Query()
	q.Add("q", city)
	q.Add("key", "9a11b5e85eea44169b9223804252611")

	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, ErrMakingRequestWeatherAPI
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrWeatherAPI
	}

	var response Response
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
