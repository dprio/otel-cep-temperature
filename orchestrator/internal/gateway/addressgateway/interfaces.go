package addressgateway

import (
	"context"

	"github.com/dprio/otel-cep-temperature/orchestrator/internal/domain/address"
	"github.com/dprio/otel-cep-temperature/orchestrator/internal/infrastructure/httpclient/viacephttpclient"
)

type (
	Client interface {
		GetAddress(ctx context.Context, cep string) (*viacephttpclient.Response, error)
	}

	Gateway interface {
		GetAddressByZipCode(ctx context.Context, zipCode address.ZipCode) (*address.Address, error)
	}
)
