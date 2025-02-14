package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB struct
type MongoDB struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// ConnectMongoDB initializes MongoDB connection
func ConnectMongoDB(uri, dbName string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to connect to MongoDB: %w", err)
	}

	db := client.Database(dbName)
	log.Println("✅ Connected to MongoDB!")

	return &MongoDB{Client: client, DB: db}, nil
}