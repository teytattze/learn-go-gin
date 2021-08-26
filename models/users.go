package models

import (
	"fmt"

	"github.com/teytattze/learn-go-gin/db/blog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName,omitempty" bson:"first_name,omitempty"`
	LastName  string             `json:"lastName,omitempty" bson:"last_name,omitempty"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
}

func usersCollection() *mongo.Collection {
	return blog.Collection(blog.USERS)
}

func GetAllUsers() ([]User, error) {
	cursor, err := usersCollection().Find(blog.Ctx(), bson.M{})
	if err != nil {
		return []User{}, err
	}
	defer cursor.Close(blog.Ctx())

	var users []User
	for cursor.Next(blog.Ctx()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return []User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserById(id primitive.ObjectID) (User, error) {
	result := usersCollection().FindOne(blog.Ctx(), bson.M{"_id": id})
	if err := result.Err(); err != nil {
		return User{}, err
	}

	var user User
	if err := result.Decode(&user); err != nil {
		return User{}, err
	}

	return user, nil
}

func CreateUser(userData interface{}) (*mongo.InsertOneResult, error) {
	result, err := usersCollection().InsertOne(blog.Ctx(), userData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}
