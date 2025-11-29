package addressgateway

import (
	"context"

	"github.com/dprio/cep-temperature/internal/domain/address"
	"github.com/dprio/cep-temperature/internal/infrastructure/httpclient/viacephttpclient"
)

type (
	Client interface {
		GetAddress(ctx context.Context, cep string) (*viacephttpclient.Response, error)
	}

	Gateway interface {
		GetAddressByZipCode(ctx context.Context, zipCode address.ZipCode) (*address.Address, error)
	}
)
