package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"learn_go/pkg/setting"
)

var db *mongo.Database
var client *mongo.Client
var userCol *mongo.Collection

// Setup connect database mongodb
func Setup() {
	// Set client options
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("%s://%s",
		setting.DatabaseSetting.Type,
		setting.DatabaseSetting.Host,
	))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("mongo 连接失败：", err) // log.Fatal 输出后会中断程序的执行
	}

	db = client.Database(setting.DatabaseSetting.Name)

	userCol = db.Collection("user")

	fmt.Println("Connected to MongoDB!")
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
