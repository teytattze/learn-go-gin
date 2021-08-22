package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title,omitempty" bson:"title,omitempty"`
	Content string             `json:"content,omitempty" bson:"content,omitempty"`
	Author  string             `json:"author,omitempty" bson:"author,omitempty"`
}

func GetAllPosts() ([]Post, error) {
	cursor, err := db.Collection(POSTS).Find(ctx, bson.M{})
	if err != nil {
		return []Post{}, err
	}
	defer cursor.Close(ctx)

	var posts []Post
	for cursor.Next(ctx) {
		var post Post
		if err = cursor.Decode(&post); err != nil {
			return []Post{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostById(id primitive.ObjectID) (Post, error) {
	result := db.Collection(POSTS).FindOne(ctx, bson.M{"_id": id})

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
	result, err := db.Collection(POSTS).InsertOne(ctx, postData)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdatePost(id primitive.ObjectID, postData interface{}) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": postData}

	result, err := db.Collection(POSTS).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, err
}
