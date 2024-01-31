package Models

import (
	"context"
	"simple-api/app/pkg/constants"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	productId primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	name      string             `json:"name" bson:"name"`
	price     float32            `json:"price" bson:"price"`
	trademark string             `json:"trademark" bson:"trademark"`
}

func (product *Product) CreateUser(db *mongo.Client) error {
	userCollection := db.Database(constants.DatabaseName).Collection(constants.ProductCollection)
	result, err := userCollection.InsertOne(context.TODO(), product)

	if err != nil {
		return err
	}

	product.productId = result.InsertedID.(primitive.ObjectID)
	return nil
}
