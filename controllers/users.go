package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/teytattze/learn-go-gin/models"
	"github.com/teytattze/learn-go-gin/pkg/ex"
	"github.com/teytattze/learn-go-gin/pkg/status"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.AbortWithStatusJSON(ex.BadRequest("", ""))
		return
	}
	c.JSON(status.SUCCESS, users)
}

func GetUserById(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(ex.BadRequest("", ""))
		return
	}

	user, err := models.GetUserById(id)
	if err != nil {
		c.AbortWithStatusJSON(ex.NotFound("", ""))
		return
	}

	c.JSON(status.SUCCESS, user)
}

type CreateUserDto struct {
	FirstName string `form:"firstName" bson:"first_name" binding:"required"`
	LastName  string `form:"lastName" bson:"last_name" binding:"required"`
	Username  string `form:"username" bson:"username" binding:"required"`
	Email     string `form:"email" bson:"email" binding:"required,email"`
	Password  string `form:"password" bson:"password" binding:"required"`
}

type CreateUserResult struct {
	Id interface{} `json:"_id"`
}

func CreateUser(c *gin.Context) {
	var userData *CreateUserDto
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.AbortWithStatusJSON(ex.BadRequest("", ""))
		return
	}

	user, err := models.CreateUser(userData)
	if err != nil {
		c.AbortWithStatusJSON(ex.BadRequest("", ""))
		return
	}

	userResult := &CreateUserResult{Id: user.InsertedID}

	c.JSON(status.SUCCESS, userResult)
}
