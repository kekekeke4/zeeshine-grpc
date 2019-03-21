package middleware

import (
	"errors"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func handleServerPainc(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, rerr error) {
	defer func() {
		if err := recover(); err != nil {
			err = errors.New(fmt.Sprintf("%v", err))
			fmt.Println(err)
		}
	}()

	return handler(ctx, req)
}

func hanlderServerStreamPainc(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	defer func() {
		if err := recover(); err != nil {
			err = errors.New(fmt.Sprintf("%v", err))
			fmt.Println(err)
		}
	}()

	return handler(srv, ss)
}
