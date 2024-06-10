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
	r.HandleFunc("/v1/user/{id}", controllers.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/v1/user", controllers.CreateUser).Methods(http.MethodPost)

	//login
	r.HandleFunc("/v1/login", controllers.GetSesion).Methods(http.MethodPost)

	//projects
	r.HandleFunc("/v1/projects", controllers.GetAllprojects).Methods(http.MethodGet)
	r.HandleFunc("/v1/projects", controllers.CreateProject).Methods(http.MethodPost)

	//tests
	r.HandleFunc("/v1/test/{id}", controllers.GetTestById).Methods(http.MethodGet)
	r.HandleFunc("/v1/test/{id}", controllers.UpdateTest).Methods(http.MethodPut)
	r.HandleFunc("/v1/test", controllers.CreateTest).Methods(http.MethodPost)

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
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT")

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
