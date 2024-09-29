package api

import (
	"context"
)

type RegistryClient interface {
	Register(ctx context.Context, serviceId string, serviceAddress string) error
	Deregister(ctx context.Context, serviceId string) error
	Discover(ctx context.Context, serviceId string) (string, error)
	// HealthCheck(ctx context.Context, serviceName string, serviceAddress string, servicePort string)
}
