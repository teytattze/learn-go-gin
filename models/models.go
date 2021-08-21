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

var client *mongo.Client

// Initialize databases connection
func Setup() {
	mongoConfig := config.MongoConfig
	mongoUri := "mongodb+srv://" + mongoConfig.Username + ":" + mongoConfig.Password + mongoConfig.Uri

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic("Error...while creating mongo instance...")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		panic("Error...while connecting to mongo databases...")
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic("Error...while pinging mongo databases...")
	}

	fmt.Println("Connected to MongoDB successfully!")
}
