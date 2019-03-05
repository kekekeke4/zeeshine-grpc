package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
	"log"
)

type consulBuilder struct {
	address     string
	client      *api.Client
	serviceName string
}

func NewConsulBuilder(address string) resolver.Builder {
	config := api.DefaultConfig()
	config.Address = address
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal("LearnGrpc: create consul client error", err.Error())
		return nil
	}
	return &consulBuilder{address: address, client: client}
}

func (cb *consulBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	cb.serviceName = target.Endpoint

	adds, serviceConfig, err := cb.resolve()
	if err != nil {
		return nil, err
	}
	cc.NewAddress(adds)
	cc.NewServiceConfig(serviceConfig)

	consulResolver := NewConsulResolver(&cc, cb, opts)
	consulResolver.wg.Add(1)
	go consulResolver.watcher()

	return consulResolver, nil
}

func (cb *consulBuilder) resolve() ([]resolver.Address, string, error) {

	serviceEntries, _, err := cb.client.Health().Service(cb.serviceName, "", true, &api.QueryOptions{})
	if err != nil {
		return nil, "", err
	}

	adds := make([]resolver.Address, 0)
	for _, serviceEntry := range serviceEntries {
		address := resolver.Address{Addr: fmt.Sprintf("%s:%d", serviceEntry.Service.Address, serviceEntry.Service.Port)}
		adds = append(adds, address)
	}
	return adds, "", nil
}

func (cb *consulBuilder) Scheme() string {
	return "consul"
}
