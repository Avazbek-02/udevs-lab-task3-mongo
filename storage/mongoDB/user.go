package mongodb

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/google/uuid"
	pb "app/genprotos/user"
)

type UserRepo struct {
	db *mongo.Database
}

func NewUserRepo(d *mongo.Database) *UserRepo {
	return &UserRepo{db: d}
}

func (repo *UserRepo) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	coll := repo.db.Collection("users")
	id := uuid.NewString()
	_, err := coll.InsertOne(ctx, bson.M{
		"id":        id,
		"firstName": req.FirstName,
		"lastName":  req.LastName,
		"email":     req.Email,
		"age":       req.Age,
	})
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return &pb.CreateUserResponse{
			Success: false,
			Message: "Failed to create user",
		}, err
	}
	return &pb.CreateUserResponse{
		Success: true,
		Message: "User created successfully",
	}, nil
}

func (repo *UserRepo) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserResponse, error) {
	coll := repo.db.Collection("users")
	var user pb.User
	err := coll.FindOne(ctx, bson.M{"id": req.Id}).Decode(&user)
	if err != nil {
		log.Printf("Failed to get user: %v", err)
		return nil, err
	}
	return &pb.GetUserResponse{
		User: &user,
	}, nil
}

func (repo *UserRepo) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	coll := repo.db.Collection("users")
	update := bson.M{
		"$set": bson.M{
			"firstName": req.FirstName,
			"lastName":  req.LastName,
			"email":     req.Email,
			"age":       req.Age,
		},
	}
	_, err := coll.UpdateOne(ctx, bson.M{"id": req.Id}, update)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return &pb.UpdateUserResponse{
			Success: false,
			Message: "Failed to update user",
		}, err
	}
	return &pb.UpdateUserResponse{
		Success: true,
		Message: "User updated successfully",
	}, nil
}

func (repo *UserRepo) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	coll := repo.db.Collection("users")
	_, err := coll.DeleteOne(ctx, bson.M{"id": req.Id})
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		return &pb.DeleteUserResponse{
			Success: false,
			Message: "Failed to delete user",
		}, err
	}
	return &pb.DeleteUserResponse{
		Success: true,
		Message: "User deleted successfully",
	}, nil
}

func (repo *UserRepo) GetAllUser(ctx context.Context, req *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error) {
	coll := repo.db.Collection("users")
	var users []*pb.User

	filter := bson.M{}
	if req.User.FirstName != "" {
		filter["firstName"] = req.User.FirstName
	}
	if req.User.LastName != "" {
		filter["lastName"] = req.User.LastName
	}
	if req.User.Email != "" {
		filter["email"] = req.User.Email
	}
	if req.User.Age != 0 {
		filter["age"] = req.User.Age
	}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		log.Printf("Failed to get users with filters %v: %v", filter, err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user pb.User
		if err := cursor.Decode(&user); err != nil {
			log.Printf("Failed to decode user: %v", err)
			return nil, err
		}

		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error iterating through users: %v", err)
		return nil, err
	}

	return &pb.GetAllUserResponse{
		User: users, 
	}, nil
}
