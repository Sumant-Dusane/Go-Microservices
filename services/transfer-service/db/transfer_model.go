package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Transfers struct {
	Id     string `json:"_id"`
	Amount string `json:"amt"`
}

var TransferCollection *mongo.Collection

func Get() ([]Transfers, error) {
	results := []Transfers{}
	cursor, err := TransferCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &results)

	if err != nil {
		return nil, err
	}

	return results, nil
}

func GetById(id string) (*Transfers, error) {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var transfer = Transfers{}

	err = TransferCollection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&transfer)

	if err != nil {
		return nil, err
	}

	return &transfer, nil
}
