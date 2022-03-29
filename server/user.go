package main

import (
	"context"
	"log"

	"github.com/SIB61/Go-gRPC/gen"
)
type Server struct {

	gen.UnimplementedUserServiceServer
}
var Users []*gen.User

func (s *Server) CreateAccount(ctx context.Context, user *gen.User) (*gen.Status, error) {
	Users=append(Users,user)
	log.Printf("user created successfully: %v" , Users)
	return &gen.Status{Status: gen.StatusType_SUCCESS}, nil
}


func (s *Server) Login(ctx context.Context,user *gen.User) (*gen.Status,error){
	fl := false

	for _,u := range Users {
		if user.GetUserName()==u.GetUserName() && user.GetPassword()==u.GetPassword() {
           fl=true
		   break
		}
	}

	if fl {
		log.Printf("Loged in successfully with username : %s",user.GetUserName())
		return &gen.Status{Status: gen.StatusType_SUCCESS}, nil
	} else {
		log.Printf("Loged in failed with username : %s",user.GetUserName())
		return &gen.Status{Status: gen.StatusType_FAILED}, nil
	}
}


func (s *Server) Logout(ctx context.Context,user *gen.User) (*gen.Status,error){
	log.Printf("Logout successfully from user : %s", user.GetUserName())
	return &gen.Status{Status: gen.StatusType_SUCCESS}, nil
}


func (s *Server) DeleteAccount(ctx context.Context,user *gen.User) (*gen.Status,error){
	log.Printf("%s deleted successfully",user.GetPassword())
	return &gen.Status{Status: gen.StatusType_SUCCESS}, nil
}



