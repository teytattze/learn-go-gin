package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teytattze/learn-go-gin/models"
	"github.com/teytattze/learn-go-gin/pkg/config"
	"github.com/teytattze/learn-go-gin/routers"
)

func init() {
	config.Setup()
	models.Setup()
}

func main() {
	gin.SetMode(config.AppConfig.Server.Mode)
	r := routers.InitRouter()
	r.Run(":" + config.AppConfig.Server.Port)
}
