package models

import (
	"context"
	"errors"

	"github.com/Ivan2001otp/REST-API-GO-lang/constants"
	"go.mongodb.org/mongo-driver/bson"
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

func FetchUserById(db *mongo.Client, id string) (*User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, errors.New("Not a valid Id")
	}

	var u User
	userCollection := db.Database(constants.DBName).Collection(constants.UserCollection)

	err = userCollection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&u)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func DeleteUserById(db *mongo.Client, id string) (string, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return "", errors.New("Not a valid id")
	}

	userCollection := db.Database(constants.DBName).Collection(constants.UserCollection)

	result, err := userCollection.DeleteOne(context.TODO(), bson.M{"_id": objectId})

	if err != nil {
		return "", err
	}

	if result.DeletedCount > 0 {
		return "document deleted successfully", nil
	}

	return "Document not found", errors.New("document not found while deleting")

}

func (u *User) UpdateModelById(db *mongo.Client, id string) (*User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, errors.New("Invalid id of user while update operation")
	}

	u.Id = objectId

	updatePtr := bson.M{}

	if u.Name != "" {
		updatePtr["name"] = u.Name
	}
	if u.Email != "" {
		updatePtr["email"] = u.Email
	}

	if u.MobileNumber != "" {
		updatePtr["mobileNumber"] = u.MobileNumber
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": updatePtr,
	}

	userCollection := db.Database(constants.DBName).Collection(constants.UserCollection)

	_, err = userCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return nil, err
	}

	res := userCollection.FindOne(context.TODO(), filter)
	res.Decode(&u)
	return u, nil
}
