package main

import (
	"jobtrack-api/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()

	// Config CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowedOrigins: []string{"http://frontend:80", "http://localhost:3000"}, // URL frontend React
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization", "X-Requested-With"},
	})
	r.Use(corsOptions.Handler)

	// Responde manualmente às requisições OPTIONS (preflight)
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization, X-Requested-With")
		w.WriteHeader(http.StatusOK)
	})

	r.HandleFunc("/", handlers.HandleRoot)
	r.HandleFunc("/applications", handlers.ListApplications).Methods("GET")
	r.HandleFunc("/applications", handlers.CreateApplication).Methods("POST")
	r.HandleFunc("/applications", handlers.UpdateStatus).Methods("PATCH")
	r.HandleFunc("/applications/{id}", handlers.DeleteApplication).Methods("DELETE")

	return r
}
