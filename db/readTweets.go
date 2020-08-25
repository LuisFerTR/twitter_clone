package db

import (
	"context"
	"log"
	"time"

	"github.com/LuisFerTR/twitter_clone/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadTweets reads profile tweets
func ReadTweets(ID string, page int64) ([]*models.ReturnTweets, bool){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter_clone")
	col := db.Collection("tweet")

	var results []*models.ReturnTweets

	condition := bson.M{
		"userid": ID,
	}

	opts := options.Find()
	opts.SetLimit(20)
	opts.SetSort(bson.D{{Key:"date", Value:-1}})
	opts.SetSkip((page - 1)*20)

	cursor, err := col.Find(ctx, condition, opts)

	if err != nil {
		log.Fatal(err.Error())
		return nil, false
	}

	for cursor.Next(context.TODO()) {
		var data models.ReturnTweets
		err := cursor.Decode(&data)
		if err != nil {
			return nil, false
		}

		results = append(results, &data)
	}

	return results, true
}