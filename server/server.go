package server

import (
	"fmt"
	"net"
	"time"

	"github.com/kekekeke4/zeeshine-grpc/discovery/consul"
	"google.golang.org/grpc"
)

type GrpcServiceServer interface {
	Register(server *grpc.Server) error
}

// Server is the host
type GrpcServer struct {
	// opts        *ServerOptions
	port        int
	grpcSvr     *grpc.Server
	serviceName string
}

func (s *GrpcServer) RegisterConsul(consulAddr string, serviceName string) *GrpcServer {
	s.serviceName = serviceName
	opts := &consul.ConsulRegisterOptions{
		ConsulAddress:                  consulAddr,
		ServiceName:                    serviceName,
		Tags:                           []string{},
		ServicePort:                    s.port,
		DeregisterCriticalServiceAfter: time.Duration(1) * time.Minute,
		Interval:                       time.Duration(10) * time.Second,
	}
	consul.RegisterConsulHealth(s.grpcSvr)
	register := consul.NewConsulRegister(opts)
	if err := register.Register(); err != nil {
		panic(err)
	}
	return s
}

// Serve 服务
func (s *GrpcServer) Serve() *GrpcServer {
	addr := fmt.Sprintf("0.0.0.0:%v", s.port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	if err := s.grpcSvr.Serve(ln); err != nil {
		panic(err)
	}
	return s
}
