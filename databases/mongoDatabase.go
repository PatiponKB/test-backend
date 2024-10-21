package databases

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongoDB() (*mongo.Client) {
	url := "mongodb://localhost:2000"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil
	}

	if err := client.Ping(ctx, readpref.Primary());
		 err != nil {
		return nil
	}

	fmt.Println("Connected to MongoDB!")
	return client
}