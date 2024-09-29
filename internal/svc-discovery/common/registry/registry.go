package registry

import (
	"context"
	"fmt"
	regapi "foodordering-svc/internal/svc-discovery/common/api"
	consul "foodordering-svc/internal/svc-discovery/consul"
)

type Registry struct {
	RegClient regapi.RegistryClient
}

func NewRegistry(ctx context.Context, regAddr string, regProvider string) (*Registry, error) {

	switch regProvider {
	case "consul":

		regClient, err := consul.NewConsulRegistry(ctx, regAddr)

		return &Registry{regClient}, err
	}

	return nil, fmt.Errorf("No implementation found for the gGiven registration provider")

}
