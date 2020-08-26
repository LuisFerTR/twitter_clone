package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/LuisFerTR/twitter_clone/middlew"
	"github.com/LuisFerTR/twitter_clone/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers sets my port, handler and gets ready to listen requests
// and call server
func Handlers() {
	// Controlará los status que devuelvan las URL
	router := mux.NewRouter()

	router.HandleFunc("/signup", middlew.CheckDB(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.LogIn)).Methods("POST")
	router.HandleFunc("/see-profile", middlew.CheckDB(middlew.ValidateJWT(routers.SeeProfile))).Methods("GET")
	router.HandleFunc("/modify-profile", middlew.CheckDB(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidateJWT(routers.WriteTweet))).Methods("POST")
	router.HandleFunc("/read-tweets", middlew.CheckDB(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/remove-tweet", middlew.CheckDB(middlew.ValidateJWT(routers.RemoveTweet))).Methods("DELETE")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	// Que IPs pueden acceder a nuestra app
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
