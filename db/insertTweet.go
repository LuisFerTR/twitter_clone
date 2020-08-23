package db

import (
	"context"
	"time"

	"github.com/LuisFerTR/twitter_clone/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertTweet writes the tweet in database
func InsertTweet(t models.WriteTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter_clone")
	col := db.Collection("tweet")

	data := bson.M{
		"userid": t.UserID,
		"message": t.Message,
		"date": t.Date,
	}

	result, err := col.InsertOne(ctx, data)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.Hex(), true, nil
}