package handlers

import (
	"encoding/json"
	"fmt"
	"jobtrack/jobtrack-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the JobTrack API!!"))
}
func ListApplications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Applications)
}
func CreateApplication(w http.ResponseWriter, r *http.Request) {
	var newApplication models.Application
	err := json.NewDecoder(r.Body).Decode(&newApplication)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newApplication.ID = len(models.Applications) + 1
	models.Applications = append(models.Applications, newApplication)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newApplication)
}

func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	var updated models.ApplicationUpdate
	err := json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, application := range models.Applications {
		if application.ID == updated.ID {
			models.Applications[i].Status = updated.Status

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(models.Applications[i])
			return
		}
	}
	http.Error(w, "No application was found with this ID", http.StatusNotFound)
}
func DeleteApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, application := range models.Applications {
		if application.ID == id {
			models.Applications = append(models.Applications[:i], models.Applications[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Application %d deleted", id)
			return
		}
	}
	http.Error(w, "No application was found with this ID", http.StatusNotFound)
}
