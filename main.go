package main

// import (
// 	"github.com/kekekeke4/zeeshine-grpc/middleware"
// 	"github.com/kekekeke4/zeeshine-grpc/server"
// 	"google.golang.org/grpc"
// )

import (
	"github.com/kekekeke4/zeeshine-grpc/client"
	"github.com/kekekeke4/zeeshine-grpc/discovery/consul"
)

func main() {
	// sb := server.NewServerBuilder()
	// sv := sb.Build(func(so *server.ServerOptions) {
	// 	so.Initialize(5250, func(mco *middleware.MiddlewareServerChainOptions) (grpc.UnaryServerInterceptor, grpc.StreamServerInterceptor) {
	// 		return mco.UsePaincMiddleware().
	// 			UseOAuthMiddleware().
	// 			BuildMiddleware()
	// 	}).
	// 		UseServices(func(grpcServer *grpc.Server) {
	// 		}).
	// 		UseConsulRegister("127.0.0.1:8500", "order_service", make([]string, 0, 1))
	// })
	// sv.Serve()

	consul.InitConsulResolver()
	client, err := client.NewGrpcClientUseConsul("consul", "192.168.2.103:8500")
	if err != nil {
		panic("sasfasfd")
	}

	_, e := client.GetConn("road-base")
	if e != nil {
		panic(e)
	}
}
