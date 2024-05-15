package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://anhhuy:%40nhHuy1010@cluster0.fmkkfsr.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}
	fmt.Println("Connected to MongoDB!")

	///////////////////////////////////////////////
	colection := client.Database("thunghiem").Collection("thu")
	document := bson.M{
		"name":  "tran bao anh huy",
		"age":   100,
		"email": "mail1@gmail.com",
	}
	_, err = colection.InsertOne(context.Background(), document)
	if err != nil {
		log.Fatal(err)
	}
}
