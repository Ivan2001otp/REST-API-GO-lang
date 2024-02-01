package models

import (
	"context"

	"github.com/Ivan2001otp/REST-API-GO-lang/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	Email        string             `json:"email" bson:"email"`
	MobileNumber string             `json:"mobileNumber" bson:"mobileNumber"`
}

func (u *User) ModelUser(db *mongo.Client) error {

	userCollection := db.Database(constants.DBName).Collection(constants.UserCollection)

	result, err := userCollection.InsertOne(context.TODO(), u)

	if err != nil {
		return err
	}

	u.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}
