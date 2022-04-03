package main

import (
	"context"
	"flag"
	//"fmt"
	"log"
	"net"

	"github.com/SIB61/Go-gRPC/pb"
	"github.com/jackc/pgx/v4"

	//"github.com/SIB61/Go-gRPC/server/db"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:8088")
	if err != nil {
		log.Fatal(err)
	}
   
	 connection,err:=pgx.Connect(context.Background(),"user=postgres database=go_grpc")
     if err!=nil{
		 log.Fatal("pgx.connect failed")
	 }
	defer connection.Close(context.Background())
	
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &Server{connection: connection})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("grpc error", err)
	}else{
		log.Println("listening at ", lis.Addr())
	}
}
