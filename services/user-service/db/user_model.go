package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id       string `json:"_id"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var UserCollection *mongo.Collection

func Get() ([]User, error) {
	results := []User{}
	cursor, err := UserCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &results)

	if err != nil {
		return nil, err
	}

	return results, nil
}

func GetById(id string) (*User, error) {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var user = User{}

	err = UserCollection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
