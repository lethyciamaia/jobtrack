package main

import (
	"jobtrack-api/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HandleRoot)
	r.HandleFunc("/applications", handlers.ListApplications).Methods("GET")
	r.HandleFunc("/applications", handlers.CreateApplication).Methods("POST")
	r.HandleFunc("/applications", handlers.UpdateStatus).Methods("PATCH")
	r.HandleFunc("/applications/{id}", handlers.DeleteApplication).Methods("DELETE")

	// Config CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://frontend:80", "http://localhost:3000"}, // URL frontend React
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})
	r.Use(corsOptions.Handler)

	return r
}
