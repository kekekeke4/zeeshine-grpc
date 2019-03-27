package consul

import (
	"google.golang.org/grpc/resolver"
)

type consulBuilder struct {
}

// impl interface
func (cb *consulBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	addr := target.Authority //fmt.Sprintf("%s%s", host, port)
	name := target.Endpoint
	cr, err := NewConsulResolver(addr, name)
	if err != nil {
		return nil, err
	}

	cr.disableServiceConfig = opts.DisableServiceConfig
	cr.cc = cc
	cr.wg.Add(1)
	go cr.watcher()
	return cr, nil
}

// impl interface
func (cb *consulBuilder) Scheme() string {
	return "consul"
}
