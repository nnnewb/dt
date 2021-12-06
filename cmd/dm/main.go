package main

import (
	"log"
	"net"

	"github.com/nnnewb/dt/internal/svc/dm"
	"github.com/nnnewb/dt/pkg/models"
	"github.com/nnnewb/dt/pkg/pb"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(mysql:3306)/dm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// sync table schema every time
	db.AutoMigrate(&models.GlobalTx{}, &models.SubTx{})

	s := grpc.NewServer()
	svc := &dm.DMService{DB: db}
	pb.RegisterDMServer(s, svc)

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
