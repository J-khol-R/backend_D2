package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/J-khol-R/backend_D2/controllers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	r := mux.NewRouter()

	// db.MigrateModels()

	//User
	r.HandleFunc("/v1/user/{id}", controllers.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/v1/user/{id}", controllers.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/v1/user", controllers.CreateUser).Methods(http.MethodPost)

	//login
	r.HandleFunc("/v1/login", controllers.GetSesion).Methods(http.MethodGet)

	//offers
	r.HandleFunc("/v1/offers", controllers.GetAllOffers).Methods(http.MethodGet)
	r.HandleFunc("/v1/offers", controllers.CreateOffer).Methods(http.MethodPost)
	r.HandleFunc("/v1/offers/{id}", controllers.GetOffer).Methods(http.MethodGet)
	r.HandleFunc("/v1/offers/{id}", controllers.UpdateOffer).Methods(http.MethodPut)
	r.HandleFunc("/v1/offers/{id}", controllers.DeleteOffer).Methods(http.MethodDelete)
	r.HandleFunc("/v1/offers/created/{idUser}", controllers.GetCreatedUserOffers).Methods(http.MethodGet)
	r.HandleFunc("/v1/offers/postulated/{idUser}", controllers.GetPostutatedUserOffers).Methods(http.MethodGet)

	//postulations
	r.HandleFunc("/v1/postulations/{idUser}", controllers.GetUserPostulations).Methods(http.MethodGet)
	r.HandleFunc("/v1/postulations", controllers.CreatePostulation).Methods(http.MethodPost)
	r.HandleFunc("/v1/postulations/{idPostulation}/{idUser}", controllers.DeletePostulation).Methods(http.MethodDelete)

	//califications
	r.HandleFunc("/v1/califications/{idUser}", controllers.GetUserCalification).Methods(http.MethodGet)
	r.HandleFunc("/v1/califications", controllers.CreateCalification).Methods(http.MethodPost)

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"PUT", "GET", "POST", "DELETE"},
	})

	handler := corsOptions.Handler(r)

	direccion := ":8081"

	servidor := &http.Server{
		Handler:      handler,
		Addr:         direccion,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Escuchando en %s. Presiona CTRL + C para salir", servidor.Addr)
	log.Fatal(servidor.ListenAndServe())
}
