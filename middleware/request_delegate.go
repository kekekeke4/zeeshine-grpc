package middleware

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// StreamRequestDelegate 流请求委托
type ServerStreamRequestDelegate func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error

// UnaryRequestDelegate 请求委托
type ServerUnaryRequestDelegate func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error)

// ClientStreamRequestDelegate
type ClientStreamRequestDelegate func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error)

// ClientUnaryRequestDelegate
type ClientUnaryRequestDelegate func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error
