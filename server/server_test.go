package server

// import(
// 	"github.com/kekekeke4/zeeshine-grpc/middleware"
// 	"google.golang.org/grpc"
// )

// func Sever_Test()  {
// 	sb:= NewServerBuilder()
// 	sv:= sb.Build(func(so *ServerOptions){
// 		so.Initialize(5150,func(mco *middleware.MiddlewareChainOptions) (grpc.UnaryServerInterceptor,grpc.StreamServerInterceptor){
// 			return	mco.UsePaincMiddleware().
// 			            UseOAuthMiddleware().
// 			            Build()
// 		}).
// 		UseServices(func(grpcServer *grpc.Server){
// 		}).
// 		UseConsulRegister("addr","serviceName",nil)
// 	})
// 	sv.Serve()
// }
