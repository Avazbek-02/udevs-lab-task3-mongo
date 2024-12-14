package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	pb "app/genprotos/user"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new user
// @Tags user
// @Description Create a new user with the provided information
// @Accept json
// @Produce json
// @Param request body pb.CreateUserRequest true "Create User Request"
// @Success 200 {object} pb.CreateUserResponse "User successfully created"
// @Failure 400 {object} string "Invalid request"
// @Router /user/create [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var req pb.CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	res, err := h.stg.CreateUser(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	fmt.Println(1234)
	c.JSON(http.StatusOK, res)
}

// @Summary Get user by ID
// @Tags user
// @Description Retrieve a user by their ID
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} pb.GetUserResponse "User found"
// @Failure 400 {object} string "Invalid request"
// @Failure 404 {object} string "User not found"
// @Router /user/get/{id} [get]
func (h *Handler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	req := &pb.GetUserByIDRequest{Id: id}
	res, err := h.stg.GetUserByID(context.Background(), req)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Update an existing user
// @Tags user
// @Description Update an existing user with new details
// @Accept json
// @Produce json
// @Param request body pb.UpdateUserRequest true "Update User Request"
// @Success 200 {object} pb.UpdateUserResponse "User successfully updated"
// @Failure 400 {object} string "Invalid request"
// @Failure 404 {object} string "User not found"
// @Router /user/update [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	var req pb.UpdateUserRequest
	if err := c.BindJSON(&req); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	res, err := h.stg.UpdateUser(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Delete a user by ID
// @Tags user
// @Description Delete a user by their ID
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} pb.DeleteUserResponse "User successfully deleted"
// @Failure 400 {object} string "Invalid request"
// @Failure 404 {object} string "User not found"
// @Router /user/delete/{id} [post]
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	req := &pb.DeleteUserRequest{Id: id}
	res, err := h.stg.DeleteUser(context.Background(), req)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Get all users with optional filters
// @Tags user
// @Description Retrieve a list of users based on the provided filters. If no filter is provided, all users are returned.
// @Accept json
// @Produce json
// @Param firstName query string false "Filter by first name"
// @Param lastName query string false "Filter by last name"
// @Param email query string false "Filter by email"
// @Param age query int false "Filter by age"
// @Success 200 {object} pb.GetAllUserResponse "Users found"
// @Failure 400 {object} string "Invalid request"
// @Router /user/get-all [get]
func (h *Handler) GetAllUser(c *gin.Context) {
	var req pb.GetAllUserRequest
	req.User = &pb.User{}
	req.User.FirstName = c.DefaultQuery("firstName", "")
	req.User.LastName = c.DefaultQuery("lastName", "")
	req.User.Email = c.DefaultQuery("email", "")

	age, err := strconv.Atoi(c.DefaultQuery("age", "0"))
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid age parameter"})
		return
	}
	req.User.Age = int32(age) 

	res, err := h.stg.GetAllUser(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	c.JSON(http.StatusOK, res)
}