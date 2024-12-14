package server

import (
	pb "app/genprotos/user"
	"app/service"
	mongodb "app/storage/mongoDB"
	"log"
	"net"

	"google.golang.org/grpc"
)

func Server() {
	db, err := mongodb.DbConnection()

	if err != nil {
		log.Fatal("Error with dbconnection", err)
		return
	}

	newServer := grpc.NewServer()
	pb.RegisterUserServiceServer(newServer, service.NewUserService(db))
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	err = newServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
