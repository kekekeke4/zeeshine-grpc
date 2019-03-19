package main

import(
	"github.com/kekekeke4/zeeshine-grpc/middleware"
	"github.com/kekekeke4/zeeshine-grpc/server"
	"google.golang.org/grpc"
)

func main()  {
	sb:= server.NewServerBuilder()
		sv:= sb.Build(func(so *server.ServerOptions){
			so.Initialize(5150,func(mco *middleware.MiddlewareChainOptions) (grpc.UnaryServerInterceptor,grpc.StreamServerInterceptor){
				return	mco.UsePaincMiddleware().
							UseOAuthMiddleware().
							BuildMiddleware()
			}).
			UseServices(func(grpcServer *grpc.Server){
			}).
			UseConsulRegister("addr","serviceName",nil)
		})
		sv.Serve()
}