package consul

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type consulHealth struct {
}

func (h *consulHealth) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *consulHealth) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
}

func RegisterConsulHealth(server *grpc.Server) {
	grpc_health_v1.RegisterHealthServer(server, &consulHealth{})
}
