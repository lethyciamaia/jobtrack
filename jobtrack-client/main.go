package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Creating applications:")
	newApplication := Application{EntrepriseName: "ENSTA Paris", Description: "Cybersecurity", CreatedAt: time.Now().Truncate(24 * time.Hour), Status: Sended}
	createApplication(newApplication)
	newApplication = Application{EntrepriseName: "Telecom Paris", Description: "IA", CreatedAt: time.Now().Truncate(24 * time.Hour), Status: Sended}
	createApplication(newApplication)
	fmt.Println("Update Status:")
	updated := ApplicationUpdate{ID: 3, Status: Interview}
	updateStatus(updated)
	fmt.Println("Listing all:")
	listApplications()
	fmt.Println("Deleting application 4:")
	deleteApplication(4)
	fmt.Println("Listing all:")
	listApplications()
}
