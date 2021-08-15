package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teytattze/learn-go-gin/errors"
	"github.com/teytattze/learn-go-gin/models"
	"github.com/teytattze/learn-go-gin/pkg/ex"
	"github.com/teytattze/learn-go-gin/pkg/status"
)

func GetAllPosts(c *gin.Context) {
	c.JSON(status.SUCCESS, models.GetAllPosts())
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")

	postId, err := strconv.Atoi(id)

	if err != nil {
		c.AbortWithStatusJSON(
			ex.BadRequestException(
				errors.ERROR_GET_POST_FAIL, errors.PostsErrorMsg[errors.ERROR_GET_POST_FAIL],
			),
		)
		return
	}

	post := models.GetPostById(postId)

	if post == (models.Post{}) {
		c.AbortWithStatusJSON(
			ex.NotFoundException(
				errors.ERROR_GET_POST_FAIL, errors.PostsErrorMsg[errors.ERROR_GET_POST_FAIL],
			),
		)
		return
	}

	c.JSON(status.SUCCESS, post)
}
