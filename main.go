package main

import (
	"github.com/kekekeke4/zeeshine-grpc/server"
)

// import (
// 	"github.com/kekekeke4/zeeshine-grpc/client"
// 	"github.com/kekekeke4/zeeshine-grpc/discovery/consul"
// )

func main() {
	sb := server.NewServerBuilder()
	sv := sb.ForPort(5350).
		UserPanicMiddleware().
		Build()
	sv.RegisterConsul("192.168.2.103:8500", "road-test").
		Serve()
}
