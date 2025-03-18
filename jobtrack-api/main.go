package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("JobTrack API in Golang")
	r := setupRouter()
	log.Fatal(http.ListenAndServe(":8081", r))
}
