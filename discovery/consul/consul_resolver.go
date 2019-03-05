package consul

import (
	"context"
	"google.golang.org/grpc/resolver"
	"log"
	"sync"
	"time"
)

type consulResolver struct {
	clientConn           *resolver.ClientConn
	consulBuilder        *consulBuilder
	t                    *time.Ticker
	wg                   sync.WaitGroup
	rn                   chan struct{}
	ctx                  context.Context
	cancel               context.CancelFunc
	disableServiceConfig bool
}

func NewConsulResolver(cc *resolver.ClientConn, cb *consulBuilder, opts resolver.BuildOption) *consulResolver {
	ctx, cancel := context.WithCancel(context.Background())
	return &consulResolver{
		clientConn:           cc,
		consulBuilder:        cb,
		t:                    time.NewTicker(time.Second),
		ctx:                  ctx,
		cancel:               cancel,
		disableServiceConfig: opts.DisableServiceConfig}
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

		adds, serviceConfig, err := cr.consulBuilder.resolve()
		if err != nil {
			log.Fatal("query service entries error:", err.Error())
		}

		(*cr.clientConn).NewAddress(adds)
		(*cr.clientConn).NewServiceConfig(serviceConfig)

		// self by kez
		// conn := (*cr.clientConn).(*consulClientConn)
		// conn.NewAddress(adds)
		// conn.NewServiceConfig(serviceConfig)
	}
}

func (cr *consulResolver) Scheme() string {
	return cr.consulBuilder.Scheme()
}

func (cr *consulResolver) ResolveNow(rno resolver.ResolveNowOption) {
	select {
	case cr.rn <- struct{}{}:
	default:
	}
}

func (cr *consulResolver) Close() {
	cr.cancel()
	cr.wg.Wait()
	cr.t.Stop()
}

func GenerateAndRegisterConsulResolver(address string, serviceName string) (schema string, err error) {
	builder := NewConsulBuilder(address)
	target := resolver.Target{Scheme: builder.Scheme(), Endpoint: serviceName}
	_, err = builder.Build(target, NewConsulClientConn(), resolver.BuildOption{})
	if err != nil {
		return builder.Scheme(), err
	}
	resolver.Register(builder)
	schema = builder.Scheme()
	return
}
