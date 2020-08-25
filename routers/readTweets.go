package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LuisFerTR/twitter_clone/db"
)

func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID parameter must be provided", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Page parameter must be provided", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "Page parameter must be a positive integer greater than zero", http.StatusBadRequest)
		return
	}

	pag := int64(page)
	response, ok := db.ReadTweets(ID, pag)

	if !ok {
		http.Error(w, "An error has ocurred reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(&response)
}