package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ivan-aguiar/phoenix/middleware"
	"github.com/ivan-aguiar/phoenix/routers"
	"github.com/rs/cors"
)

/*Handlers setea el puerto, el handler y pone en escucha el server*/
func Handlers() {

	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckDatabase(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
