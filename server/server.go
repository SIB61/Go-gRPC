package main

import (
	"context"
	"log"
//	"strings"

	"github.com/SIB61/Go-gRPC/gen"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
	gen.UnimplementedUserServiceServer
}

var Users []*gen.User

func (s *Server) CreateAccount(ctx context.Context, user *gen.User) (*gen.Status, error) {
	e:=s.db.Create(&gen.User{Email: user.Email,Password: user.Password}).Error
	if e!=nil{
		log.Fatal("new user created with email: "+user.Email)
		return &gen.Status{Status: gen.StatusType_SUCCESS}, nil
	}else{
		log.Fatal("falied")
		return &gen.Status{Status: gen.StatusType_FAILED}, e
	}
}

func (s *Server) Login(ctx context.Context, user *gen.User) (*gen.Status, error) {
	return &gen.Status{}, nil
}

func (s *Server) Logout(ctx context.Context, user *gen.User) (*gen.Status, error) {
	return &gen.Status{}, nil
}

func (s *Server) DeleteAccount(ctx context.Context, user *gen.User) (*gen.Status, error) {
	return &gen.Status{}, nil
}
