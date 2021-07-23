package main

import (
	"fmt"
	entities "github.com/florian74/randomise/entities"
	"github.com/florian74/randomise/random"
	grpc "google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	host := ":8082"
	srv := grpc.NewServer()
	impl := random.GetInstance()
	entities.RegisterRandomiseServer(srv, impl)
	list, err := net.Listen("tcp", host)
	if err != nil {
		panic("could not listen on " + host)
	}
	if err := srv.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Printf("server stop")

}
