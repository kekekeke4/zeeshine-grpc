package consul

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
)

var (
	errMissingAddr   = errors.New("consul resolver: missing address")
	errAddrMisMatch  = errors.New("consul resolver: invalied uri")
	errEndsWithColon = errors.New("consul resolver: missing port after port-separator colon")
	regexConsul, _   = regexp.Compile("^([A-z0-9.]+)(:[0-9]{1,5})?/([A-z_]+)$")
)

func InitConsulResolver() {
	resolver.Register(&consulBuilder{})
}

type consulResolver struct {
	address              string
	wg                   sync.WaitGroup
	cc                   resolver.ClientConn
	serviceName          string
	disableServiceConfig bool
	t                    *time.Ticker
	rn                   chan struct{}
	ctx                  context.Context
	cancel               context.CancelFunc
	consulClient         *api.Client
}

func NewConsulResolver(addr string, serviceName string) (*consulResolver, error) {
	config := api.DefaultConfig()
	config.Address = addr
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	cr := &consulResolver{
		address:      addr,
		serviceName:  serviceName,
		t:            time.NewTicker(time.Second),
		ctx:          ctx,
		cancel:       cancel,
		consulClient: client,
	}
	return cr, nil
}

// impl interface
func (cr *consulResolver) ResolveNow(opt resolver.ResolveNowOption) {
	select {
	case cr.rn <- struct{}{}:
	default:
	}
}

// impl interface
func (cr *consulResolver) Close() {
	cr.cancel()
	cr.wg.Wait()
	cr.t.Stop()
}

func (cr *consulResolver) watcher() {
	cr.wg.Done()
	for {
		select {
		case <-cr.ctx.Done():
			return
		case <-cr.rn:
		case <-cr.t.C:
		}

		addrs, serviceConfig, err := cr.resolve()
		if err != nil {
			fmt.Println("query service entries error:", err.Error())
		}
		cr.cc.NewAddress(addrs)
		cr.cc.NewServiceConfig(serviceConfig)
	}
}

func (cr *consulResolver) resolve() ([]resolver.Address, string, error) {
	services, _, err := cr.consulClient.Health().Service(cr.serviceName, "", true, &api.QueryOptions{})
	if err != nil {
		return nil, "", err
	}

	addrs := make([]resolver.Address, 0)
	for _, service := range services {
		addr := resolver.Address{Addr: fmt.Sprintf("%s:%d", service.Service.Address, service.Service.Port)}
		addrs = append(addrs, addr)
	}

	return addrs, "", nil
}
