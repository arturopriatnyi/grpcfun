package server

import (
	"context"
	"fmt"

	"github.com/imarrche/grpcfun/pkg/api"
)

type Server struct{}

func (s *Server) Greet(ctx context.Context, req *api.GreetRequest) (*api.GreetResponse, error) {
	return &api.GreetResponse{Greeting: fmt.Sprintf("Hello, %s!", req.GetName())}, nil
}
