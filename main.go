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

	// password, err := security.HashPassword("contraseña")
	// if err != nil {
	// 	return
	// }
	// score := models.Score{
	// 	Nro_estrellas: "3",
	// 	Comentarios: "excelente",
	// 	Fecha: "hoy",
	// 	Calificador: "",
	// }

	// // Utilizar el método Create para insertar el usuario en la base de datos
	// if err := db.Conn.Create(&newUser).Error; err != nil {
	// 	panic(err)
	// }

	// nuevaOferta := models.Offer{
	// 	IDServicio:  1, // Reemplaza con el ID del servicio real
	// 	IDOfertante: 4, // Reemplaza con el ID del ofertante real
	// 	Estado:      "Pendiente",
	// }

	// // Crear la oferta en la base de datos
	// db.Conn.Create(&nuevaOferta)

	// // Verificar errores
	// if db.Conn.Error != nil {
	// 	// Manejar el error, por ejemplo, imprimirlo
	// 	fmt.Println(db.Conn.Error)
	// 	return
	// }

	//User
	r.HandleFunc("/v1/user/{id}", controllers.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/v1/user/{id}", controllers.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/v1/user", controllers.CreateUser).Methods(http.MethodPost)

	//login
	r.HandleFunc("/v1/login", controllers.GetSesion).Methods(http.MethodPost)

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
