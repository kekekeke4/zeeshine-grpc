package middleware

import(
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

type MiddlewareChainOptionsAction func(mco *MiddlewareChainOptions) (grpc.UnaryServerInterceptor,grpc.StreamServerInterceptor , error)

type MiddlewareChainOptions struct{
	unaryInters []grpc.UnaryServerInterceptor
	streamInters []grpc.StreamServerInterceptor
}

func NewMiddlewareChainOptions() *MiddlewareChainOptions{
	return &MiddlewareChainOptions{
		unaryInters:make([]grpc.UnaryServerInterceptor,0,10),
	}
}

func (mo *MiddlewareChainOptions)AddUnary(delegate UnaryRequestDelegate) (*MiddlewareChainOptions, error){
	inter:=grpc.UnaryServerInterceptor(delegate)
	mo.unaryInters = append(mo.unaryInters,inter)
	return mo,nil
}

func (mo *MiddlewareChainOptions)AddStream(delegate StreamRequestDelegate) (*MiddlewareChainOptions, error){
	inter:=grpc.StreamServerInterceptor(delegate)
	mo.streamInters = append(mo.streamInters,inter)
	return mo,nil
}

func (mo *MiddlewareChainOptions)UsePaincMiddleware()(*MiddlewareChainOptions,error){
	return mo.AddUnary(handlePainc)
}

func (mo *MiddlewareChainOptions)UseOAuthMiddleware()(*MiddlewareChainOptions,error){
	return mo.AddUnary(handleOAuth)
}

func(mo *MiddlewareChainOptions) Build()(grpc.UnaryServerInterceptor,grpc.StreamServerInterceptor ,error){
	var unary grpc.UnaryServerInterceptor
	var stream grpc.StreamServerInterceptor 

	if len(mo.unaryInters)>0 {
		unary = grpc_middleware.ChainUnaryServer(mo.unaryInters...) 
	}

	if len(mo.streamInters)>0 {
		stream = grpc_middleware.ChainStreamServer(mo.streamInters...) 
	}

	return unary,stream,nil
}
