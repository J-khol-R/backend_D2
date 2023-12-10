package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/J-khol-R/backend_D2/controllers"
	"github.com/J-khol-R/backend_D2/db"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	db.MigrateModels()

	//User
	//obtener el usuario con el id
	r.HandleFunc("/v1/user/{id}", controllers.GetUser).Methods(http.MethodGet)
	//actualizar usuario
	r.HandleFunc("/v1/user/{id}", controllers.UpdateUser).Methods(http.MethodPut)
	//crear usuario
	r.HandleFunc("/v1/user", controllers.CreateUser).Methods(http.MethodPost) //vale

	//login
	//obtener el token de sesion y el id del usuario
	r.HandleFunc("/v1/login", controllers.GetSesion).Methods(http.MethodPost) //vale

	//offers
	//obtener todas las ofertas disponibles (una oferta ya no es disponible cuando acepto a alguien que me trabaje)
	r.HandleFunc("/v1/offers", controllers.GetAllOffers).Methods(http.MethodGet)
	//obtener una oferta por su id
	r.HandleFunc("/v1/offers/{id}", controllers.GetOffer).Methods(http.MethodGet)
	//borrar una oferta por su id
	r.HandleFunc("/v1/offers/{id}", controllers.DeleteOffer).Methods(http.MethodDelete)
	//obtener las ofertas que ha creado un usuario
	r.HandleFunc("/v1/offers/created/{idUser}", controllers.GetCreatedUserOffers).Methods(http.MethodGet)
	//obtener las ofertas a las que se ha postulado un usuario
	r.HandleFunc("/v1/offers/postulated/{idUser}", controllers.GetPostutatedUserOffers).Methods(http.MethodGet)
	//crear una oferta
	r.HandleFunc("/v1/offer", controllers.CreateOffer).Methods(http.MethodPost)
	//actualizar una oferta
	r.HandleFunc("/v1/offer/{id}", controllers.UpdateOffer).Methods(http.MethodPut) //vale

	//postulations
	//obtener las postulaciones asociadas a una oferta
	r.HandleFunc("/v1/postulations/{idOffer}", controllers.GetOfferPostulations).Methods(http.MethodGet)
	//crear una postulacion
	r.HandleFunc("/v1/postulations", controllers.CreatePostulation).Methods(http.MethodPost)
	//eliminar una postulacion (cuando alguien ya no quiere estar postulado a esa oferta)
	r.HandleFunc("/v1/postulations/{idPostulation}", controllers.DeletePostulation).Methods(http.MethodDelete)

	//califications
	//obtener las calificaciones que s ele han hecho a un usuario
	r.HandleFunc("/v1/califications/{idUser}", controllers.GetUserCalifications).Methods(http.MethodGet)
	//crear una calificacion
	r.HandleFunc("/v1/califications", controllers.CreateCalification).Methods(http.MethodPost) //listo

	// corsOptions := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedMethods: []string{"PUT", "GET", "POST", "DELETE"},
	// })

	handler := corsMiddleware(r)

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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Permitir todas las origenes
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Permitir los métodos HTTP especificados
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE")
		w.Header().Set("Access-Control-Allow-Methods", "PUT")

		// Permitir los encabezados HTTP especificados
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Permitir el envío de cookies
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Si la solicitud es una solicitud OPTIONS, simplemente respondemos con los encabezados CORS sin continuar con la cadena de middleware
		if r.Method == "OPTIONS" {
			return
		}

		// Continuar con la cadena de middleware
		next.ServeHTTP(w, r)
	})
}
