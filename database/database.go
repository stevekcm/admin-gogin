package database

import (
	"context"
	"log"

	config "github.com/stevekcm/admin-gogin/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Database, error) {
	config.InitEnvConfigs()

	log.Println("Connecting to MongoDB...", config.EnvConfigs.MongoUri, config.EnvConfigs.MongoUri)

	// Set up the MongoDB client options
	clientOptions := options.Client().ApplyURI(config.EnvConfigs.MongoUri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, err
	}

	// ping to check connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	db := client.Database(config.EnvConfigs.MongoName)

	return db, nil
}
