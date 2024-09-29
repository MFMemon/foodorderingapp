package consul

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	// regapi "foodordering-svc/pkg/svc-discovery/common/api"

	// "log"

	consul "github.com/hashicorp/consul/api"
)

type ConsulRegistry struct {
	client *consul.Client
}

func NewConsulRegistry(ctx context.Context, servaddr string) (*ConsulRegistry, error) {
	conf := consul.DefaultConfig()
	conf.Address = servaddr

	client, err := consul.NewClient(conf)

	return &ConsulRegistry{client}, err
}

func (r *ConsulRegistry) Register(ctx context.Context, serviceId string, serviceAddress string) error {

	svcHostAndPort := strings.Split(serviceAddress, ":")
	svcHost := svcHostAndPort[0]
	svcPort, _ := strconv.Atoi(svcHostAndPort[1])

	err := r.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		Name:    serviceId,
		ID:      serviceId,
		Address: svcHost,
		Port:    svcPort,
	})

	return err
}

func (r *ConsulRegistry) Deregister(ctx context.Context, serviceId string) error {
	return nil
}

func (r *ConsulRegistry) Discover(ctx context.Context, serviceId string) (string, error) {

	status, svcMeta, _ := r.client.Agent().AgentHealthServiceByID(serviceId)

	if status == "critical" {
		return "", fmt.Errorf("Error. Service is either down or not in healthy condition")
	}

	svcAddr := svcMeta.Service.Address + fmt.Sprintf(":%v", svcMeta.Service.Port)

	return svcAddr, nil
}
