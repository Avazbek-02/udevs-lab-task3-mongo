package service

import (
	pb "app/genprotos/user"
	mongodb "app/storage/mongoDB"
	"context"
)

type UserService struct {
	Storage *mongodb.Storage
	pb.UnimplementedUserServiceServer
}

func NewUserService(repo *mongodb.Storage) *UserService {
	return &UserService{Storage: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res, err := s.Storage.UserI.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserResponse, error) {
	res, err := s.Storage.UserI.GetUserByID(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	res, err := s.Storage.UserI.UpdateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	res, err := s.Storage.UserI.DeleteUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) GetAllUser(ctx context.Context, req *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error){
	res, err := s.Storage.UserI.GetAllUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}