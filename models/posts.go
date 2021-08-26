package models

import (
	"github.com/teytattze/learn-go-gin/db/blog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
	Id      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title,omitempty" bson:"title,omitempty"`
	Content string             `json:"content,omitempty" bson:"content,omitempty"`
	Author  string             `json:"author,omitempty" bson:"author,omitempty"`
}

func postsCollection() *mongo.Collection {
	return blog.Collection(blog.POSTS)
}

func GetAllPosts() ([]Post, error) {
	cursor, err := postsCollection().Find(blog.Ctx(), bson.M{})
	if err != nil {
		return []Post{}, err
	}
	defer cursor.Close(blog.Ctx())

	var posts []Post
	for cursor.Next(blog.Ctx()) {
		var post Post
		if err = cursor.Decode(&post); err != nil {
			return []Post{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostById(id primitive.ObjectID) (Post, error) {
	result := postsCollection().FindOne(blog.Ctx(), bson.M{"_id": id})
	if err := result.Err(); err != nil {
		return Post{}, err
	}

	var post Post
	if err := result.Decode(&post); err != nil {
		return Post{}, err
	}

	return post, nil
}

func CreatePost(postData interface{}) (*mongo.InsertOneResult, error) {
	result, err := postsCollection().InsertOne(blog.Ctx(), postData)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdatePost(id primitive.ObjectID, postData interface{}) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": postData}

	result, err := postsCollection().UpdateOne(blog.Ctx(), filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeletePost(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}

	result, err := postsCollection().DeleteOne(blog.Ctx(), filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
