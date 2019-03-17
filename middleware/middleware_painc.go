package middleware

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// // PaincMiddleware is the deal painc when call grpc service
// var PaincMiddleware RequestDelegate

// func init() {
// 	PaincMiddleware = handlePainc
// }

func handlePainc (ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, rerr error)   {
	defer func() {
		if err := recover(); err != nil {
			resp = nil
			rerr = errors.New(fmt.Sprintf("%v", err)) 
			fmt.Println(err)
		}
	}()

	return handler(ctx, req)
}