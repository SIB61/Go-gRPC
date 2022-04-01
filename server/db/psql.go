package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	//dsn := "host=localhost user=gorm password=gorm dbname=grpc_go port=9920 sslmode=disable TimeZone=Asia/Dhaka"
	dsn := "user=postgres dbname=go_grpc"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err!=nil {
	   fmt.Println(err)
	}else {
		fmt.Println("connected to postgres")
	}
  return db
}


