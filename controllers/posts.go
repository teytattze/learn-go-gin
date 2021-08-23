package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/teytattze/learn-go-gin/errors"
	"github.com/teytattze/learn-go-gin/models"
	"github.com/teytattze/learn-go-gin/pkg/ex"
	"github.com/teytattze/learn-go-gin/pkg/status"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Handler to get all posts
func GetAllPosts(c *gin.Context) {
	posts, err := models.GetAllPosts()
	if err != nil {
		c.AbortWithStatusJSON(ex.InternalServer("", ""))
		return
	}

	c.JSON(status.SUCCESS, posts)
}

// Hanlder to get post by id
func GetPostById(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(ex.BadRequest(errors.ERROR_GET_POST_FAIL, errors.PostsErrorMsg[errors.ERROR_GET_POST_FAIL]))
		return
	}

	post, err := models.GetPostById(id)
	if err != nil {
		c.AbortWithStatusJSON(ex.InternalServer("", ""))
		return
	}

	if post == (models.Post{}) {
		c.AbortWithStatusJSON(ex.NotFound(errors.ERROR_GET_POST_FAIL, errors.PostsErrorMsg[errors.ERROR_GET_POST_FAIL]))
		return
	}

	c.JSON(status.SUCCESS, post)
}

type CreatePostDto struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
	Author  string `form:"author" binding:"required"`
}

type CreatePostResult struct {
	Id interface{} `json:"_id"`
}

// Handler to create a new post
func CreatePost(c *gin.Context) {
	var postData *CreatePostDto
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.AbortWithStatusJSON(ex.BadRequest("", "Bad Request"))
		return
	}

	post, err := models.CreatePost(postData)
	if err != nil {
		c.AbortWithStatusJSON(ex.BadRequest("", "Bad Request"))
		return
	}

	postResult := &CreatePostResult{Id: post.InsertedID}

	c.JSON(status.SUCCESS, postResult)
}

type UpdatePostDto struct {
	Title   string `form:"title" bson:",omitempty"`
	Content string `form:"content" bson:",omitempty"`
	Author  string `form:"author" bson:",omitempty"`
}

// Handler to update single post
func UpdatePost(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(ex.BadRequest("", ""))
		return
	}

	var postData *UpdatePostDto
	if err = c.ShouldBindJSON(&postData); err != nil {
		c.AbortWithStatusJSON(ex.BadRequest("", ""))
		return
	}

	_, err = models.UpdatePost(id, postData)
	if err != nil {
		c.AbortWithStatusJSON(ex.BadRequest("", ""))
		return
	}

	c.JSON(status.SUCCESS, map[string]interface{}{
		"status": status.SUCCESS, "message": "Update successfully", "_id": id.Hex(),
	})
}

// Handler to delete single post
func DeletePost(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(ex.BadRequest("", ""))
		return
	}

	_, err = models.DeletePost(id)
	if err != nil {
		c.AbortWithStatusJSON(ex.BadRequest("", ""))
		return
	}

	c.JSON(status.SUCCESS, map[string]interface{}{
		"status": status.SUCCESS, "message": "Deleted successfully", "_id": id.Hex(),
	})
}
