package storage

import (
	"context"
	pb "app/genprotos/user"
)

type StorageI interface {
	User() UserI
}

type UserI interface {
	CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserResponse, error)
	UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
	GetAllUser(ctx context.Context, req *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error)
}
