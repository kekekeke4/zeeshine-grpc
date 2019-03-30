package server

import (
	"github.com/kekekeke4/zeeshine-grpc/middleware"
	"google.golang.org/grpc"
)

type ServerBuilder struct {
	port        int
	middlewares *middleware.MiddlewareServerChainOptions
	services    []GrpcServiceServer
}

// // ServerBuildAction 服务构建行为函数
// type ServerBuildAction func(so *ServerOptions)

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{
		middlewares: middleware.NewMiddlewareServerChainOptions(),
		services:    make([]GrpcServiceServer, 0, 10),
	}
}

// func (sb *ServerBuilder) Build(action ServerBuildAction) *GrpcServer {
// 	opts := new(ServerOptions)
// 	action(opts)
// 	server := new(GrpcServer)
// 	server.opts = opts
// 	return server
// }

func (sb *ServerBuilder) ForPort(port int) *ServerBuilder {
	sb.port = port
	return sb
}

func (sb *ServerBuilder) UserPanicMiddleware() *ServerBuilder {
	sb.middlewares.UsePaincMiddleware()
	return sb
}

func (sb *ServerBuilder) AddUnaryMiddleware(interceptor grpc.UnaryServerInterceptor) *ServerBuilder {
	sb.middlewares.AddUnaryMiddleware(interceptor)
	return sb
}

func (sb *ServerBuilder) AddStreamMiddleware(interceptor grpc.StreamServerInterceptor) *ServerBuilder {
	sb.middlewares.AddStreamMiddleware(interceptor)
	return sb
}

func (sb *ServerBuilder) AddService(service GrpcServiceServer) *ServerBuilder {
	sb.services = append(sb.services, service)
	return sb
}

func (sb *ServerBuilder) Build() *GrpcServer {
	unaryInter, streamInter := sb.middlewares.BuildMiddleware()
	grpcOpts := make([]grpc.ServerOption, 0, 2)
	if unaryInter != nil {
		grpcOpts = append(grpcOpts, grpc.UnaryInterceptor(unaryInter))
	}

	if streamInter != nil {
		grpcOpts = append(grpcOpts, grpc.StreamInterceptor(streamInter))
	}

	grpcSvr := grpc.NewServer(grpcOpts...)
	for _, s := range sb.services {
		s.Register(grpcSvr)
	}

	return &GrpcServer{
		port:    sb.port,
		grpcSvr: grpcSvr,
	}
}
