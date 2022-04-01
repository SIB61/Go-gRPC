package main

import (
	"context"
	"fmt"
	"log"

	"github.com/SIB61/Go-gRPC/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	client := gen.NewUserServiceClient(conn)

	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Create an account")
		fmt.Println("2. Login into an existing account")
		var i int
		fmt.Scan(&i)
		if i == 1 {
			var (
				email    string
				password string
			)
			fmt.Scan(&email, &password)
			user := gen.User{Email: email, Password: password}
			status, err := client.CreateAccount(context.Background(), &user)
            if err!=nil{
				fmt.Println(err)
			}

			if status.GetStatus() == gen.StatusType_SUCCESS {
				log.Println("User created successfully")
			}
		} else if i == 2 {

		}
	}

}
