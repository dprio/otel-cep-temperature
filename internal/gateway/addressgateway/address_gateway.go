package addressgateway

import (
	"context"
	"strings"

	"github.com/dprio/otel-cep-temperature/internal/domain/address"
)

type addressgateway struct {
	client Client
}

func New(client Client) Gateway {
	return &addressgateway{
		client: client,
	}
}

func (g *addressgateway) GetAddressByZipCode(ctx context.Context, zipCode address.ZipCode) (*address.Address, error) {
	resp, err := g.client.GetAddress(ctx, zipCode.Value())
	if err != nil {
		return nil, err
	}

	return address.New(strings.ReplaceAll(resp.ZipCode, "-", ""), resp.City)
}
