package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	PostsRouter(api)

	return r
}
