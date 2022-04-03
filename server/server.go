package main

import (

	//	"strings"

	"context"
	"fmt"
	"log"

	"github.com/SIB61/Go-gRPC/pb"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	//"gorm.io/gorm"
)

type Server struct {
	connection *pgx.Conn
	pb.UnimplementedUserServiceServer
}

func (server *Server) Register(ctx context.Context, user *pb.User) (*pb.Response, error) {
	fmt.Println(user.Email)

	email := user.Email
	password, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal("password hashing failed")
	}

	tx, err := server.connection.Begin(context.Background())
	if err != nil {
		log.Fatal("connection failed")
	}
	_, err = tx.Exec(context.Background(), "insert into users(email,password) values ($1, $2)", email, password)
	if err != nil {
		log.Fatal("tx.Exec failed")
		
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Fatal("tx.commmit failed")
		err = grpc.Errorf(codes.Unknown, "Unknown Error")
	} 
	return &pb.Response{}, err
	
}

func (s *Server) Login(ctx context.Context, user *pb.User) (*pb.Response, error) {
	_,hPassword:=s.getUserCredentials(user.Email)
	var err error
	if !CheckPasswordHash(user.Password, hPassword) {
		err=status.Error(codes.Unauthenticated,"AUTHENTICATION FAILED")
		return &pb.Response{}, err
		//return &pb.Response{},err
	}
	return &pb.Response{}, err
}

func (s *Server) DeleteAccount(ctx context.Context, user *pb.User) (*pb.Response, error) {
    _,hPassword:=s.getUserCredentials(user.Email)
	var err error
	if !CheckPasswordHash(user.Password, hPassword) {
		err=status.Error(codes.Unauthenticated,"AUTHENTICATION FAILED")
		return &pb.Response{}, err
	}
	tx, err := s.connection.Begin(context.Background())
	if err != nil {
		log.Fatal("tx.begine failed failed")
	}
	_,err=tx.Exec(context.Background(),"DELETE FROM users WHERE email=$1",user.Email)
	if err!=nil{
		err=status.Error(codes.Unknown,"FAILED")
	}
	err = tx.Commit(context.Background())

	return &pb.Response{}, err
}

func (s *Server)getUserCredentials(mail string)(string, string){
	var (
		email     string
		hPassword string
	)
	row := s.connection.QueryRow(context.Background(), "SELECT * FROM users WHERE email=$1",mail)
	row.Scan(&email, &hPassword)
    return email,hPassword
}
