package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/imarrche/grpcfun/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("invalid number of arguments")
	}

	name := flag.Arg(0)

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewGreeterClient(conn)

	res, err := c.Greet(context.Background(), &api.GreetRequest{Name: name})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.GetGreeting())
}
