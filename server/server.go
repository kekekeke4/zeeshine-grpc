package server

import (
	"fmt"
	"net"
)

// Server is the host
type Server struct {
	opts *ServerOptions
	// grpcServer *grpc.Server
}

// Serve 服务
func (s *Server) Serve() error {
	port := s.opts.GetServerPort()
	addr := fmt.Sprintf("0.0.0.0:%v", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	grpcServer := s.opts.grpcServer
	if err := grpcServer.Serve(ln); err != nil {
		panic(err)
	}
	return nil
}
