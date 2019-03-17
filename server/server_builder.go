package server

import(
	"github.com/zeeshine/grpc/middleware"
	"google.golang.org/grpc"
)

type ServerBuilder struct{
	grpcServer *grpc.Server
}

// ServerBuildAction 服务构建行为函数
type ServerBuildAction func(so *ServerOptions) error

func NewServerBuilder() *ServerBuilder{
	return &ServerBuilder{}
}

func (sb *ServerBuilder)Initialize(action middleware.MiddlewareChainOptionsAction) (*ServerBuilder,error){
	opts:=middleware.NewMiddlewareChainOptions()
	 unaryInter,streamInter, err:= action(opts)
	if err != nil {
		return sb,err
	}

	grpcOpts:=make([]grpc.ServerOption, 0,2)

	if unaryInter !=nil {
		grpcOpts = append(grpcOpts,grpc.UnaryInterceptor(unaryInter))
	}

	if streamInter !=nil{
		grpcOpts = append(grpcOpts,grpc.StreamInterceptor(streamInter))
	}

	sb.grpcServer  = grpc.NewServer(grpcOpts...)
	return sb,nil
}

func (sb *ServerBuilder) Build(serverPort int, action ServerBuildAction) (*Server,error){
	opts:= &ServerOptions{
		serverPort:serverPort,
		grpcServer:sb.grpcServer,
	}

	err := action(opts)
	if err!=nil{
		return nil,err
	}

	server := new(Server)
	server.opts = opts
	server.grpcServer = sb.grpcServer
	return nil,nil	
}