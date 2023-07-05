package routes

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func DBInstance() error {
	MongoDb := "mongodb://localhost:27017/caloriesDb"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoDb))
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	Client = client
	fmt.Println("Connected to MongoDB")
	return nil
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("caloriesDb").Collection(collectionName)
	return collection
}
