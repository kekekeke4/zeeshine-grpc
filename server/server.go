package server

import (
	"net"
	"fmt"
	"google.golang.org/grpc"
)

// Server is the host 
type Server struct {
	opts *ServerOptions
	grpcServer *grpc.Server
}

// Serve 服务
func(s *Server) Serve() error{
	port:=s.opts.GetServerPort()
	addr,err:=net.Listen("tcp",fmt.Sprintf("0.0.0.0:%v", port))
	if err!=nil{
		panic(err)
	}

	if err:=s.grpcServer.Serve(addr); err!=nil{
		panic(err)
	}

	return nil
}