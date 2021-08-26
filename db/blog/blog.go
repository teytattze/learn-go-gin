package blog

import (
	"context"
	"time"

	"github.com/teytattze/learn-go-gin/db"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	BLOG_DB = "blog_db"

	USERS = "users"
	POSTS = "posts"
)

func Ctx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func Collection(c string) *mongo.Collection {
	return db.MI.Select(BLOG_DB, c)
}
