package main

import (
	"context"
	"log"

	"github.com/SIB61/Go-gRPC/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
func main(){
	conn,err := grpc.Dial(":9000",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil {
		log.Println(err)
	}
	defer conn.Close()
	client := gen.NewUserServiceClient(conn)
    status,err := client.CreateAccount(context.Background(),&gen.User{UserName: "sabit",Password: "1234"})
    if(status.GetStatus() == gen.StatusType_SUCCESS){
		log.Println("User created successfully")
	}
}
