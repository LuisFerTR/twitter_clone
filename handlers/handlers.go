package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers sets my port, handler and gets ready to listen requests
// and call server
func Handlers() {
	// Controlará los status que devuelvan las URL
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	// Que IPs pueden acceder a nuestra app
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
