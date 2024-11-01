package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(cfg *Configuration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Database.Timeout)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(cfg.Database.URI).
		SetServerAPIOptions(serverAPI).
		SetMinPoolSize(10).
		SetMaxPoolSize(100)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
		return nil, err
	}

	// Ping the database
	if err := client.Database("admin").RunCommand(ctx, map[string]string{"ping": "1"}).Err(); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to MongoDB")
	return client, nil
}
