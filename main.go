package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teytattze/learn-go-gin/routers"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")

	routers.PostsRouter(api)

	router.Run(":8080")
}
