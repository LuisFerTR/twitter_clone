package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/LuisFerTR/twitter_clone/db"
	"github.com/LuisFerTR/twitter_clone/models"
)

// WriteTweet writes tweet data in database
func WriteTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet

	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		http.Error(w, "Wrong data "+err.Error(), http.StatusBadRequest)
		return
	}

	data := models.WriteTweet{
		UserID: IDUser,
		Message: message.Message,
		Date: time.Now(),
	}

	_, status, err := db.InsertTweet(data)

	if err != nil {
		http.Error(w, "An error has occurred inserting tweet in database. Try again "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Tweet couldn't be written in database", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}