package main

import (
	"flag"
	//"fmt"
	"log"
	"net"

	"github.com/SIB61/Go-gRPC/gen"
	"github.com/SIB61/Go-gRPC/server/db"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}

	psql := db.New()
    print(psql)
    psql.AutoMigrate(&gen.User{})

	
	grpcServer := grpc.NewServer()
	gen.RegisterUserServiceServer(grpcServer, &Server{db: psql})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("grpc error", err)
	}else{
		log.Println("listening at ", lis.Addr())
	}

}
