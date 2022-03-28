package main

import (
	"context"
	"flag"
	"github.com/SIB61/Go-gRPC/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	gen.UnimplementedUserServiceServer
}

func (s *server) CreateAccount(ctx context.Context, accountDetails *gen.AccountDetails) (*gen.Status, error) {
	log.Printf("recieved messagebody from client: %s",accountDetails.Email )
	return &gen.Status{Status: gen.StatusType_SUCCESS}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	gen.RegisterUserServiceServer(grpcServer,&server{})

	log.Println("listening at ", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("grpc error", err)
	}
}
