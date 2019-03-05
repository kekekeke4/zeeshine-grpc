package consul

import (
	"google.golang.org/grpc/resolver"
)

type consulClientConn struct {
	addrs []resolver.Address
	sc    string
}

func NewConsulClientConn() resolver.ClientConn {
	return &consulClientConn{}
}

func (cc *consulClientConn) NewAddress(addrs []resolver.Address) {
	cc.addrs = addrs
}

func (cc *consulClientConn) NewServiceConfig(serviceConfig string) {
	cc.sc = serviceConfig
}
