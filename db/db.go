package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func Init() {
	uri := os.Getenv("MONGO_URI")
	mongo, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Mongo client initialized with URI: %s \n", uri)

	mongoClient = mongo
}

func GetMongo() *mongo.Client {
	if mongoClient == nil {
		panic("Mongo not initialized yet!")
	}

	return mongoClient
}
