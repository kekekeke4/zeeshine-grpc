package middleware

type RequestDelegate func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error)

type MiddlewareChain struct {
	chain []RequestDelegate
}

func (mc *MiddlewareChain) Add(delegate RequestDelegate) {
}

func (mc *MiddlewareChain) Build() {
}
