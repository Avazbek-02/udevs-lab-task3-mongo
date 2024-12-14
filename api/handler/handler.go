package handler

import (
	pb "app/genprotos/user"
)

type Handler struct {
	stg pb.UserServiceClient
}

func NewHandler(user pb.UserServiceClient) *Handler {
	return &Handler{
		stg: user,
	}
}
