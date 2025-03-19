package main

import (
	"fmt"
	"jobtrack/jobtrack-api/database"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("JobTrack API in Golang")

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error when loading .env")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	r := setupRouter()
	database.ConnectDB(dsn)

	appPort := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(appPort, r))
}
