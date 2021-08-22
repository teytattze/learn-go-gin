package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/teytattze/learn-go-gin/controllers"
)

func PostsRouter(api *gin.RouterGroup) {
	r := api.Group("/posts")

	r.GET("", controllers.GetAllPosts)
	r.GET("/:id", controllers.GetPostById)
	r.POST("/", controllers.CreatePost)
	r.PATCH("/:id", controllers.UpdatePost)
	r.DELETE("/:id")
}
