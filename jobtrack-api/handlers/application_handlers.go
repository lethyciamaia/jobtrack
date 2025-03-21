package handlers

import (
	"encoding/json"
	"fmt"
	"jobtrack-api/database"
	"jobtrack-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the JobTrack API!!"))
}
func ListApplications(w http.ResponseWriter, r *http.Request) {
	var applications []models.Application

	result := database.DB.Find(&applications)
	if result.Error != nil {
		http.Error(w, "Failed to fetch applications", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(applications)
}
func CreateApplication(w http.ResponseWriter, r *http.Request) {
	var newApplication models.Application
	err := json.NewDecoder(r.Body).Decode(&newApplication)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	database.DB.Create(&newApplication)

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

	var application models.Application
	result := database.DB.First(&application, updated.ID)
	if result.Error != nil {
		http.Error(w, "No application was found with this ID", http.StatusNotFound)
		return
	}
	application.Status = updated.Status
	database.DB.Save(&application)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(application)
}
func DeleteApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var application models.Application
	if result := database.DB.First(&application, uint(id)); result.Error != nil {
		http.Error(w, "No application was found with this ID", http.StatusNotFound)
		return
	}

	database.DB.Delete(&application)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application %d deleted", id)
}
