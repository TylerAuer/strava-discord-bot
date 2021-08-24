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

func (m Mongo) connect() (context.Context, *mongo.Client, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.uri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	return ctx, client, cancel
}

func (m Mongo) disconnect(ctx context.Context, client *mongo.Client) {
	err := client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
}

func (m Mongo) Upsert(ad ActivityDetails) {
	ctx, client, cancel := m.connect()
	defer cancel()
	defer m.disconnect(ctx, client)

	collection := os.Getenv("MONGO_COLLECTION")

	coll := client.Database("strava").Collection(collection)

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", fmt.Sprint(ad.ID)}}
	update := bson.D{
		{"$setOnInsert", bson.D{
			{"_id", fmt.Sprint(ad.ID)},
		}},
		{"$set", ad},
		{"$set", bson.D{
			{"kraftee_name", ad.krafteeWhoRecordedActivity().SafeFirstName()},
		}},
	}

	result, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		fmt.Println("Error upserting into MongoDB: " + err.Error())
	} else {
		fmt.Println("MongoDB upsert results")
		fmt.Println("_id: " + fmt.Sprint(result.UpsertedID))
		fmt.Println("Upserted: " + fmt.Sprint(result.UpsertedCount))
	}
}

func (m Mongo) Delete(ad ActivityDetails) {
	ctx, client, cancel := m.connect()
	defer cancel()
	defer m.disconnect(ctx, client)

	collection := os.Getenv("MONGO_COLLECTION")

	coll := client.Database("strava").Collection(collection)

	filter := bson.D{{"_id", fmt.Sprint(ad.ID)}}

	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println("Error deleting from MongoDB: " + err.Error())
	} else {
		fmt.Println("MongoDB delete results")
		fmt.Println("Deleted: " + fmt.Sprint(result.DeletedCount))
	}
}
