package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/teytattze/learn-go-gin/controllers"
)

func UserRouter(api *gin.RouterGroup) {
	r := api.Group("/users")

	r.GET("", controllers.GetAllUsers)
	r.GET("/:id", controllers.GetUserById)
	r.POST("/register", controllers.CreateUser)
	r.PATCH("/:id")
	r.DELETE("/:id")
}
