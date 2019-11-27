package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// SetUser SetPi
func SetUser() {
	collection := db.Collection("user")
	_, err := collection.InsertOne(context.TODO(), bson.M{"name": "abc", "password": "mmm"})
	if err != nil {
		log.Fatal(err)
	}
}
