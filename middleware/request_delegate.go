package middleware

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// StreamRequestDelegate 流请求委托
type StreamRequestDelegate func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error

// UnaryRequestDelegate 请求委托
type UnaryRequestDelegate func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error)
