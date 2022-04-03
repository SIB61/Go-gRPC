package main

import (
	"context"
	"fmt"
	"log"

	"github.com/SIB61/Go-gRPC/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial(":8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Create an account")
		fmt.Println("2. Login into your account")
		fmt.Println("3. Delete your account")
		var i int
		_, e := fmt.Scan(&i)
		if e != nil {
			fmt.Println("enter 1 or 2")
			continue
		}
		var (
			email    string
			password string
		)
		fmt.Print("ENTER YOUR EMAIL: ")
		fmt.Scan(&email)
		fmt.Print("ENTER YOUR PASSWORD: ")
		fmt.Scan(&password)
		user := pb.User{Email: email, Password: password}

		if i == 1 {
			_, err := client.Register(context.Background(), &user)

			errStatus, _ := status.FromError(err)
			if errStatus.Code() == codes.OK {
				fmt.Println("Registered successfully")
			}else{
				fmt.Println("Registration failed")
			}
		} else if i == 2 {
			_, err := client.Login(context.Background(), &user)

			errStatus, _ := status.FromError(err)
			if errStatus.Code() == codes.OK {
				fmt.Println("Logged in successfully")
			}else{
				fmt.Println("Log in failed")
			}
			fmt.Println(err)

		} else if i==3 {
			_,err:=client.DeleteAccount(context.Background(),&user)
			errStatus, _ := status.FromError(err)
			if errStatus.Code() == codes.OK {
				fmt.Println("Account deleted successfully")
			}else{
				fmt.Println("failed")
			}
			fmt.Println(err)
		} else{
			fmt.Println("enter 1 or 2")
			continue
		}
	}
}
