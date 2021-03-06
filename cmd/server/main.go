package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/imarrche/grpcfun/internal/server"
	"github.com/imarrche/grpcfun/pkg/api"
)

func main() {
	s := grpc.NewServer()
	srv := &server.Server{}

	api.RegisterGreeterServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
