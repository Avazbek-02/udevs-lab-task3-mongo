package api

import (
	"app/api/handler"
	pb "app/genprotos/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite user service
// @version 1.0
// @description user service
// @host localhost:8088
// @BasePath /
// @in header
// @name Authourization
func Engine(user pb.UserServiceClient) *gin.Engine {
	router := gin.Default()
	url := ginSwagger.URL("swagger/doc.json")

	router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r := router.Group("/user")
	h := handler.NewHandler(user)
	{
		r.POST("/create",h.CreateUser)
		r.PUT("/update",h.UpdateUser)
		r.GET("/get/:id",h.GetUserByID)
		r.POST("/delete/:id",h.DeleteUser)
		r.GET("/get-all",h.GetAllUser)
	}
	return router
}
