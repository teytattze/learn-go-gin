package db

import (
	"context"
	"fmt"
	"time"

	"github.com/teytattze/learn-go-gin/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	Ctx    context.Context
	Client *mongo.Client
}

func (m *Mongo) Select(db string, c string) *mongo.Collection {
	return m.Client.Database(db).Collection(c)
}

var MI *Mongo

func Setup() {
	// Initialize MongoDb connection
	mongoConfig := config.MongoConfig
	mongoUri := "mongodb+srv://" + mongoConfig.Username + ":" + mongoConfig.Password + mongoConfig.Uri + mongoConfig.Database + "?retryWrites=true&w=majority"

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic("Error...while creating mongo instance...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic("Error...while connecting to mongo databases...")
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic("Error...while pinging mongo databases...")
	}

	fmt.Println("Connected to MongoDB successfully!")

	MI = &Mongo{
		Ctx:    ctx,
		Client: client,
	}
}
