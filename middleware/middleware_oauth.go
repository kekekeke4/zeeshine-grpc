package middleware

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)


func handleOAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, rerr error)   {
   // do oauth...
   return handler(ctx, req)
}