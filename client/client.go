package client

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

type GrpcClient struct {
	sync.RWMutex
	clientConns  map[string]*grpc.ClientConn
	registerAddr string
	scheme       string
}

func NewGrpcClientUseConsul(scheme string, consulAddr string) (*GrpcClient, error) {
	return &GrpcClient{registerAddr: consulAddr, scheme: scheme}, nil
}

func (gc *GrpcClient) ForkConn(serviceName string) (*grpc.ClientConn, error) {
	gc.RLock()
	if conn, ok := gc.clientConns[serviceName]; ok {
		gc.RUnlock()
		return conn, nil
	}
	gc.RUnlock()

	gc.Lock()
	defer gc.Unlock()

	// schema, err := consul.GenerateAndRegisterConsulResolver(gc.registerAddr, serviceName)
	// if err != nil {
	// 	return nil, err
	// }

	// if conn, err := grpc.Dial(fmt.Sprintf("%s:///%s", schema, serviceName), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name)); err != nil {
	// 	return nil, err
	// }

	//target      = "consul://127.0.0.1:8500/helloworld"
	target := fmt.Sprintf("%s://%s/%s", gc.scheme, gc.registerAddr, serviceName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, target, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		cancel()
		return nil, err
	}

	gc.clientConns[serviceName] = conn
	return conn, nil
}

func (gc *GrpcClient) CloseConn(serviceName string) error {
	gc.Lock()
	defer gc.Unlock()
	if conn, ok := gc.clientConns[serviceName]; ok {
		delete(gc.clientConns, serviceName)
		return conn.Close()
	}
	return nil
}
