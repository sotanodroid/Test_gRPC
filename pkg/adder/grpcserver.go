package adder

import (
	"Test_gRPC/pkg/api"
	"context"
)

// GRPCServer ...
type GRPCServer struct {
}

// Add ...
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
