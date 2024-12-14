package main

import (
	"app/api"
	_ "app/docs"
	pb "app/genprotos/user"
	"app/server"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	go server.Server()

	auth, err := grpc.NewClient("user_service:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}

	client := pb.NewUserServiceClient(auth)
	r := api.Engine(client)
	r.Run(":8088")

}
//kodi run qilish uchun make run