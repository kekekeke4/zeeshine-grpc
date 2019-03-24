package server

type ServerBuilder struct {
}

// ServerBuildAction 服务构建行为函数
type ServerBuildAction func(so *ServerOptions)

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{}
}

func (sb *ServerBuilder) Build(action ServerBuildAction) *GrpcServer {
	opts := new(ServerOptions)
	action(opts)
	server := new(GrpcServer)
	server.opts = opts
	return server
}
