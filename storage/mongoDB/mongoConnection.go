package mongodb

import (
	"app/config"
	"app/storage"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Mongo *mongo.Database
	UserI storage.UserI
}

func DbConnection() (*Storage, error) {
	config:=config.Load()
	clientOptions := options.Client().ApplyURI(config.MongoDbConnection)
	fmt.Println(config.MongoDbConnection)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error: Couldn't connect to the database.", err)
	}

	fmt.Println("Connected to MongoDB!")

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("admin")
	user := NewUserRepo(db)
	return &Storage{
		Mongo: db,
		UserI: user,
		}, err
}


func (stg *Storage) User() *storage.UserI {
	if stg.UserI == nil{
		stg.UserI = NewUserRepo(stg.Mongo)
	}
	return &stg.UserI
}