package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"learn_go/pkg/setting"
)

var db *mongo.Database
var client *mongo.Client

// Setup connect database mongodb
func Setup() {
	// Set client options
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("%s://%s", setting.Databasesetting.Type, setting.Databasesetting.Host))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(setting.Databasesetting.Name)

	fmt.Println("Connected to MongoDB!")
	// log.Fatalf("Connected to MongoDB!")
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
