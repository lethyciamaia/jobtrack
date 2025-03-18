package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Creating applications:")
	newApplication := Application{EntrepriseName: "Montimage", Description: "Cybersecurity", Date: time.Now().Truncate(24 * time.Hour), Status: Sended}
	createApplication(newApplication)
	newApplication = Application{EntrepriseName: "Ensta", Description: "Cycle d'Inge", Date: time.Now().Truncate(24 * time.Hour), Status: Sended}
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
