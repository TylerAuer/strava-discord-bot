package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	username string
	password string
	uri      string
}

func initMongo() Mongo {
	username := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	uri := "mongodb+srv://" + username + ":" + password + "@cluster0.ktraa.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

	return Mongo{username, password, uri}
}

func (m Mongo) connect() (context.Context, *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.uri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	return ctx, client
}

func (m Mongo) disconnect(ctx context.Context, client *mongo.Client) {
	err := client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
}

func (m Mongo) store() {
	ctx, client := m.connect()
	defer m.disconnect(ctx, client)

	// Insert a document
	_, err := client.Database("Cluster0").Collection("customers").InsertOne(ctx, bson.M{
		"name":  "John",
		"email": "fjaskl",
	})
	if err != nil {
		panic(err)
	}
}
