package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Creating applications:")
	newApplication := Application{EntrepriseName: "ENSTA Paris", Description: "Cybersecurity", CreatedAt: time.Now(), Status: Sended}
	createApplication(newApplication)
	newApplication = Application{EntrepriseName: "Telecom Paris", Description: "IA", CreatedAt: time.Now(), Status: Sended}
	createApplication(newApplication)
	fmt.Println("Update Status:")
	updated := ApplicationUpdate{ID: 2, Status: Interview}
	updateStatus(updated)
	fmt.Println("Listing all:")
	listApplications()
	fmt.Println("Deleting application 1:")
	deleteApplication(2)
	fmt.Println("Listing all:")
	listApplications()
}
