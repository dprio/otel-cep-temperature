package temporchestratorclient

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dprio/otel-cep-temperature/input/internal/infrastructure/httpclient/client"
	"github.com/dprio/otel-cep-temperature/input/pkg/opentelemetry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

var (
	ErrMakingRequestOrchestratorAPI = errors.New("error making request to Orchestrator")
	ErrOrchestratorAPI              = errors.New("orchestrator API returned an error")
	ErrCreatingRequest              = errors.New("error creating request to  Orchestrator API")
)

type (
	Client interface {
		GetCityWeatherInformation(ctx context.Context, cep string) (*Response, error)
	}

	weatherAPIClient struct {
		httpClient client.HttpClient
		basePath   string
	}
)

func New(httpClient client.HttpClient) Client {
	basePath := "http://cep-temp-orch:8080"

	return &weatherAPIClient{
		httpClient: httpClient,
		basePath:   basePath,
	}
}

func (c *weatherAPIClient) GetCityWeatherInformation(ctx context.Context, cep string) (*Response, error) {
	ctx, span := opentelemetry.StartSpan(ctx, "TempOrchestratorClient.GetCityWeatherInformation")
	defer span.End()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.basePath+"/addresses/"+cep+"/weather/temperature", nil)
	if err != nil {
		return nil, ErrCreatingRequest
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrOrchestratorAPI
	}

	var response Response
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
