package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func insertReputation(rep Reputation) {

	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	database, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := database.Database("threatdata").Collection("reputations")

	insertResult, err := collection.InsertOne(context.TODO(), rep)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)

}

func lookupReputation(sha256 string) Reputation {

	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	// Connect to the MongoDB and return Client instance
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Access a MongoDB collection through a database
	col := client.Database("threatdata").Collection("reputations")

	var result Reputation

	// Get a MongoDB document using the FindOne() method
	err = col.FindOne(context.TODO(), bson.M{"sha256": sha256}).Decode(&result)

	if err != nil {
		fmt.Println("Error calling FindOne():", err)
		os.Exit(1)
	} else {
		fmt.Println("FindOne() result:", result)
		fmt.Println("FindOne() Rep:", result.Rep)
		fmt.Println("FindOne() Dept:", result.DateAdded)
	}

	return result

}
