package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN is connection's database object
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://luisAdmin:OlZVkz1RIsZLsM8I@twitterclone.4hdl7.mongodb.net/twitter_clone?retryWrites=true&w=majority")

// ConnectDB allows connect to database
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connected succesfully :D")

	return client
}

// CheckConnection makes a ping to database
func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)

	return err == nil
}
