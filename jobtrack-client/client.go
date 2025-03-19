package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const baseURL = "http://localhost:8081"

type Status string

const (
	Sended    = "sended"
	Interview = "interview"
	Rejected  = "rejected"
	Accepted  = "accepted"
)

type Application struct {
	ID             int       `json:"id"`
	EntrepriseName string    `json:"entreprise"`
	Description    string    `json:"description"`
	CreatedAt      time.Time `json:"created_at"`
	Status         Status    `json:"status"`
}

type ApplicationUpdate struct {
	ID     int    `json:"id"`
	Status Status `json:"status"`
}

func sendRequest(method, url string, data interface{}) (*http.Response, error) {
	var jsonData []byte
	var err error

	if data != nil {
		jsonData, err = json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("error when serializing struct: %w", err)
		}
	}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error when creating resquest: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error when making resquest: %w", err)
	}

	return response, nil
}

func listApplications() {
	response, err := sendRequest("GET", baseURL+"/applications", nil)
	http.Get(baseURL + "/applications")
	if err != nil {
		log.Println("Error when listing applications", err)
		return
	}
	defer response.Body.Close()
	var applications []Application
	if err := json.NewDecoder(response.Body).Decode(&applications); err != nil {
		log.Println("Error when decoding response", err)
		return
	}

	for _, application := range applications {
		fmt.Println(application)
	}
}
func createApplication(application Application) {
	response, err := sendRequest("POST", baseURL+"/applications", application)
	if err != nil {
		log.Println("Error when creating application", err)
		return
	}
	defer response.Body.Close()

	var newApplication Application
	if err := json.NewDecoder(response.Body).Decode(&newApplication); err != nil {
		log.Println("Error when decoding response", err)
		return
	}

	fmt.Println("New application created: ", newApplication)
}
func updateStatus(updated ApplicationUpdate) {
	response, err := sendRequest("PATCH", baseURL+"/applications", updated)
	if err != nil {
		log.Println("Error when updating status", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Server Error: %s. ", response.Status)
		return
	}

	var application Application
	if err := json.NewDecoder(response.Body).Decode(&application); err != nil {
		log.Println("Error when decoding updated application: ", err)
		return
	}
	fmt.Println("Updated application: ", application)

}
func deleteApplication(id int) {
	response, err := sendRequest("DELETE", baseURL+"/applications/"+fmt.Sprint(id), nil)
	if err != nil {
		log.Println("Error when deleting application", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s. ", response.Status)
		return
	}
	fmt.Println("Application deleted")
}
