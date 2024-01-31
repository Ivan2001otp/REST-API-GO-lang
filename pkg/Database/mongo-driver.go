package Database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

func GetMongoClient(uri string) (*mongo.Client, error) {

	if clientInstance == nil {
		clientOptions := options.Client().ApplyURI(uri)

		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			panic(err)
		}

		clientInstance = client
	}

	return clientInstance, nil
}

func GetDB() *mongo.Client {

	return clientInstance
}

func CloseMongoConnection(client *mongo.Client) bool {
	error := client.Disconnect(context.TODO())

	if error != nil {
		panic(error)
	}

	fmt.Println("Connection is terminated from MongoDB")
	return true
}
