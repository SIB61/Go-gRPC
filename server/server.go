package main

import (
	"flag"
	"github.com/SIB61/Go-gRPC/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	gen.RegisterUserServiceServer(grpcServer,&Server{})

	log.Println("listening at ", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("grpc error", err)
	}
}
