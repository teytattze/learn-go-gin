package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// r.Use(gin.Logger())
	// r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
	// 	if err, ok := err.(string); ok {
	// 		c.AbortWithStatusJSON(ex.InternalServer("", err))
	// 	}
	// 	c.AbortWithStatusJSON(ex.InternalServer("", ""))
	// }))

	api := r.Group("/api")

	PostsRouter(api)

	return r
}
