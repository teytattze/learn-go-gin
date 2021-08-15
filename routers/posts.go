package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/teytattze/learn-go-gin/controllers"
)

func PostsRouter(r *gin.RouterGroup) {
	router := r.Group("/posts")

	router.GET("", controllers.GetAllPosts)
	router.GET("/:id", controllers.GetPostById)
	router.POST("/")
	router.PATCH("/:id")
	router.DELETE("/:id")
}
