package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mClient *mongo.Client

func InitDB() {
	uri := "mongodb://mongodb:27017"

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}

	// Ping to check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("❌ MongoDB not reachable: %v", err)
	}

	log.Println("✅ Connected to MongoDB!")
	mClient = client

	TransferCollection = mClient.Database("swami_samartha").Collection("transfers")
}

func CloseDB() {
	if err := mClient.Disconnect(context.TODO()); err != nil {
		log.Fatalf("❌ Error while disconnecting MongoDB: %v", err)
	}
	log.Println("✅ MongoDB connection closed!")
}
