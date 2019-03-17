package server

import(
	"time"
	"github.com/zeeshine/grpc/discovery"
	"github.com/zeeshine/grpc/discovery/consul"
	"google.golang.org/grpc"
)

// ServerOptions is the Server options
type ServerOptions struct {
	register discovery.Register
	serverPort int
	grpcServer *grpc.Server
}

// UseConsulRegister 应用Consul注册中心
func(so *ServerOptions)UseConsulRegister(consulAddr string,serviceName string,tags []string) (*ServerOptions,error){
	regOpts:= &consul.ConsulRegisterOptions{
		ConsulAddress:consulAddr,
		ServiceName:serviceName,
		Tags:tags,
		ServicePort:so.serverPort,
		DeregisterCriticalServiceAfter:time.Duration(1) * time.Minute,
		Interval:time.Duration(10)*time.Second,
	}

	consul.RegisterConsulHealth(so.grpcServer)
	so.register = consul.NewConsulRegister(regOpts)
	so.register.Register()
	return so,nil
}

// GetServerPort 获取服务器监听端口
func(so *ServerOptions) GetServerPort() int{
	return so.serverPort
}