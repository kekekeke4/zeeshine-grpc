package middleware

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type MiddlewareServerChainOptionsAction func(mco *MiddlewareServerChainOptions) (grpc.UnaryServerInterceptor, grpc.StreamServerInterceptor)

type MiddlewareServerChainOptions struct {
	unaryInters  []grpc.UnaryServerInterceptor
	streamInters []grpc.StreamServerInterceptor
}

func NewMiddlewareServerChainOptions() *MiddlewareServerChainOptions {
	return &MiddlewareServerChainOptions{
		unaryInters:  make([]grpc.UnaryServerInterceptor, 0, 10),
		streamInters: make([]grpc.StreamServerInterceptor, 0, 10),
	}
}

func (mo *MiddlewareServerChainOptions) AddUnary(delegate ServerUnaryRequestDelegate) *MiddlewareServerChainOptions {
	inter := grpc.UnaryServerInterceptor(delegate)
	mo.unaryInters = append(mo.unaryInters, inter)
	return mo
}

func (mo *MiddlewareServerChainOptions) AddStream(delegate ServerStreamRequestDelegate) *MiddlewareServerChainOptions {
	inter := grpc.StreamServerInterceptor(delegate)
	mo.streamInters = append(mo.streamInters, inter)
	return mo
}

func (mo *MiddlewareServerChainOptions) UsePaincMiddleware() *MiddlewareServerChainOptions {
	return mo.AddUnary(handlePainc).
		AddStream(hanlderStreamPainc)
}

func (mo *MiddlewareServerChainOptions) UseOAuthMiddleware() *MiddlewareServerChainOptions {
	return mo.AddUnary(handleOAuth)
}

func (mo *MiddlewareServerChainOptions) BuildMiddleware() (grpc.UnaryServerInterceptor, grpc.StreamServerInterceptor) {
	var unary grpc.UnaryServerInterceptor
	var stream grpc.StreamServerInterceptor

	if len(mo.unaryInters) > 0 {
		unary = mo.chainUnaryServerMiddleware()
	}

	if len(mo.streamInters) > 0 {
		stream = mo.chainStreamServer()
	}

	return unary, stream
}

func (mo *MiddlewareServerChainOptions) chainUnaryServerMiddleware() grpc.UnaryServerInterceptor {
	count := len(mo.unaryInters)
	if count == 1 {
		return mo.unaryInters[0]
	}

	if count > 1 {
		lastI := count - 1
		return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			var (
				chainHandler grpc.UnaryHandler
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				if curI == lastI {
					return handler(currentCtx, currentReq)
				}
				curI++
				resp, err := mo.unaryInters[curI](currentCtx, currentReq, info, chainHandler)
				curI--
				return resp, err
			}

			return mo.unaryInters[0](ctx, req, info, chainHandler)
		}
	}

	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(ctx, req)
	}
}

func (mo *MiddlewareServerChainOptions) chainStreamServer() grpc.StreamServerInterceptor {
	n := len(mo.streamInters)
	if n > 1 {
		lastI := n - 1
		return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			var (
				chainHandler grpc.StreamHandler
				curI         int
			)

			chainHandler = func(currentSrv interface{}, currentStream grpc.ServerStream) error {
				if curI == lastI {
					return handler(currentSrv, currentStream)
				}
				curI++
				err := mo.streamInters[curI](currentSrv, currentStream, info, chainHandler)
				curI--
				return err
			}

			return mo.streamInters[0](srv, stream, info, chainHandler)
		}
	}

	if n == 1 {
		return mo.streamInters[0]
	}

	return func(srv interface{}, stream grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return handler(srv, stream)
	}
}

type MiddlewareClientChainOptionsAction func(mco *MiddlewareClientChainOptions) (grpc.UnaryClientInterceptor, grpc.StreamClientInterceptor)

type MiddlewareClientChainOptions struct {
	unaryInters  []grpc.UnaryClientInterceptor
	streamInters []grpc.StreamClientInterceptor
}

func NewMiddlewareClientChainOptions() *MiddlewareClientChainOptions {
	return &MiddlewareClientChainOptions{
		unaryInters:  make([]grpc.UnaryClientInterceptor, 0, 10),
		streamInters: make([]grpc.StreamClientInterceptor, 0, 10),
	}
}

func (mo *MiddlewareClientChainOptions) AddUnary(delegate ClientUnaryRequestDelegate) *MiddlewareClientChainOptions {
	inter := grpc.UnaryClientInterceptor(delegate)
	mo.unaryInters = append(mo.unaryInters, inter)
	return mo
}

func (mo *MiddlewareClientChainOptions) AddStream(delegate ClientStreamRequestDelegate) *MiddlewareClientChainOptions {
	inter := grpc.StreamClientInterceptor(delegate)
	mo.streamInters = append(mo.streamInters, inter)
	return mo
}
