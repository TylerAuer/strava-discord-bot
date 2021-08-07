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

func (m Mongo) Upsert() {
	ctx, client := m.connect()
	defer m.disconnect(ctx, client)

	coll := client.Database("Cluster0").Collection("customers")

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", "987654321"}}
	update := bson.D{{"$set", bson.D{{"email", "newemail@example.com"}}}}

	coll.UpdateOne(ctx, filter, update, opts)

	// Insert a document
	// _, err := client.Database("Cluster0").Collection("customers").InsertOne(ctx, bson.M{
	// 	"_id":   "123456789",
	// 	"name":  "John",
	// 	"email": "fjaskl",
	// })
	// _, err := client.Database("Cluster0").Collection("customers").UpdateOne(ctx, bson.M{
	// 	"_id":   "123456789",
	// 	"name":  "Jimmy",
	// 	"email": "Johnson",
	// })

	// _, err = client.Database("Cluster0").Collection("customers").UpdateOne(ctx, bson.M{
	// 	"_id":   "123456789",
	// 	"name":  "Jimmy",
	// 	"email": "Johnson",
	// })
	// if err != nil {
	// 	panic(err)
	// }
}
