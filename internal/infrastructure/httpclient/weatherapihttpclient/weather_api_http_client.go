package weatherapihttpclient

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dprio/cep-temperature/internal/infrastructure/httpclient"
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
		httpClient httpclient.HttpClient
		basePath   string
	}
)

func New(httpClient httpclient.HttpClient) Client {
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
