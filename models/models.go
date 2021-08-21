package models

import (
	"context"
	"fmt"
	"time"

	"github.com/teytattze/learn-go-gin/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

// Initialize databases connection
func Setup() {
	mongoConfig := config.AppConfig.Mongo
	mongoUri := "mongodb+srv://" + mongoConfig.Username + ":" + mongoConfig.Password + mongoConfig.Uri

	Client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic("Error...while creating mongo instance...")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = Client.Connect(ctx)
	if err != nil {
		panic("Error...while connecting to mongo databases...")
	}
	defer Client.Disconnect(ctx)

	err = Client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic("Error...while pinging mongo databases...")
	}

	fmt.Println("Connected to MongoDB successfully!")
}
