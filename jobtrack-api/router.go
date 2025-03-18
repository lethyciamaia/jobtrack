package main

import (
	"jobtrack/jobtrack-api/handlers"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HandleRoot)
	r.HandleFunc("/applications", handlers.ListApplications).Methods("GET")
	r.HandleFunc("/applications", handlers.CreateApplication).Methods("POST")
	r.HandleFunc("/applications", handlers.UpdateStatus).Methods("PATCH")
	r.HandleFunc("/applications/{id}", handlers.DeleteApplication).Methods("DELETE")
	// r.HandleFunc("/applications/{}", findByParam).Methods("GET")

	return r
}
