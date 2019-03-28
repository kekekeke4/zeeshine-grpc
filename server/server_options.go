package server

// import (
// 	"time"

// 	"github.com/kekekeke4/zeeshine-grpc/discovery"
// 	"github.com/kekekeke4/zeeshine-grpc/discovery/consul"
// 	"github.com/kekekeke4/zeeshine-grpc/middleware"
// 	"google.golang.org/grpc"
// )

// // ServerOptions is the Server options
// type ServerOptions struct {
// 	register    discovery.Register
// 	serverPort  int
// 	grpcServer  *grpc.Server
// 	initialized bool
// }

// // PublicServiceAction 发布服务Action
// type PublicServiceAction func(grpcServer *grpc.Server)

// // Initialize 初始化Options
// func (so *ServerOptions) Initialize(serverPort int, action middleware.MiddlewareServerChainOptionsAction) *ServerOptions {
// 	opts := middleware.NewMiddlewareServerChainOptions()
// 	unaryInter, streamInter := action(opts)
// 	grpcOpts := make([]grpc.ServerOption, 0, 2)
// 	if unaryInter != nil {
// 		grpcOpts = append(grpcOpts, grpc.UnaryInterceptor(unaryInter))
// 	}

// 	if streamInter != nil {
// 		grpcOpts = append(grpcOpts, grpc.StreamInterceptor(streamInter))
// 	}

// 	so.grpcServer = grpc.NewServer(grpcOpts...)
// 	so.serverPort = serverPort
// 	so.initialized = true
// 	return so
// }

// // UseServices 应用服务(发布远程服务)
// func (so *ServerOptions) UseServices(action PublicServiceAction) *ServerOptions {
// 	so.assertInitialized()
// 	action(so.grpcServer)
// 	return so
// }

// // UseConsulRegister 应用Consul注册中心
// func (so *ServerOptions) UseConsulRegister(consulAddr string, serviceName string, tags []string) *ServerOptions {
// 	so.assertInitialized()
// 	regOpts := &consul.ConsulRegisterOptions{
// 		ConsulAddress:                  consulAddr,
// 		ServiceName:                    serviceName,
// 		Tags:                           tags,
// 		ServicePort:                    so.serverPort,
// 		DeregisterCriticalServiceAfter: time.Duration(1) * time.Minute,
// 		Interval:                       time.Duration(10) * time.Second,
// 	}

// 	consul.RegisterConsulHealth(so.grpcServer)
// 	so.register = consul.NewConsulRegister(regOpts)
// 	if err := so.register.Register(); err != nil {
// 		panic(err)
// 	}
// 	return so
// }

// // GetServerPort 获取服务器监听端口
// func (so *ServerOptions) GetServerPort() int {
// 	return so.serverPort
// }

// func (so *ServerOptions) assertInitialized() {
// 	if !so.initialized {
// 		panic("must call initialize before")
// 	}
// }
