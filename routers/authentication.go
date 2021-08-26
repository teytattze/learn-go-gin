package routers

import (
	"github.com/gin-gonic/gin"
)

func AuthenticationRouter(api *gin.RouterGroup) {
	r := api.Group("/auth")

	r.GET("")
	r.GET("/:id")
	r.POST("/create")
	r.PATCH("/:id")
	r.DELETE("/:id")
}
