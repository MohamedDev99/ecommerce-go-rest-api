package db

import (
	"context"

	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB Setup
// Global variables
// var DB_CLIENT *mongo.Client

// Connect to MongoDB
func ConnectToMongoDB() *mongo.Client {
	// get MongoDB URI from .env file
	MONGO_DB_URI := configs.EnvMongoURI()
	if MONGO_DB_URI == "" {
		configs.Logger.Error("MONGO_DB_URI is not set ")
	}

	// Connect to MongoDB
	mongodb_client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_DB_URI))
	if err != nil {
		configs.Logger.Error("MongoDB connection error : " + err.Error())
	}

	if err := mongodb_client.Ping(context.TODO(), nil); err != nil {
		configs.Logger.Error("MongoDB connection error : " + err.Error())
	}

	// print MongoDB connection status
	// fmt.Println("MongoDB connection status:", mongodb_client.Ping(context.TODO(), nil))
	configs.Logger.Info("✅ Successfully connected to MongoDB ...")

	return mongodb_client
}

// close MongoDB connection
func CloseMongoDBConnection(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	configs.Logger.Info("MongoDB connection closed")
}

// get collection by name
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	// get MongoDB name from .env file
	MONGO_DB_NAME := configs.EnvMongoDBName()
	if MONGO_DB_NAME == "" {
		configs.Logger.Error("MONGO_DB_NAME is not set : ")
	}
	// get collection
	return client.Database(MONGO_DB_NAME).Collection(collectionName)
}
