package httpclient

import (
	"net/http"
	"time"

	"github.com/dprio/otel-cep-temperature/input/internal/infrastructure/httpclient/temporchestratorclient"
)

type HTTPClients struct {
	TempOrchestratorClient temporchestratorclient.Client
}

func New() *HTTPClients {
	return &HTTPClients{
		TempOrchestratorClient: temporchestratorclient.New(&http.Client{Timeout: time.Second * 10}),
	}

}
