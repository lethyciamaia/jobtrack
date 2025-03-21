package main

import (
	"fmt"
	"jobtrack-api/database"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("JobTrack API in Golang")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	// dsn := "host=postgres user=postgres password=futureintern031917 dbname=jobtrack port=5432 sslmode=disable"

	r := setupRouter()
	database.ConnectDB(dsn)

	goPort := fmt.Sprintf(":%s", os.Getenv("GO_PORT"))
	log.Fatal(http.ListenAndServe(goPort, r))
}
