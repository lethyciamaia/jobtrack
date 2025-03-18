package models

import (
	"time"
)

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
	Date           time.Time `json:"date"`
	Status         Status    `json:"status"`
}

type ApplicationUpdate struct {
	ID     int    `json:"id"`
	Status Status `json:"status"`
}

var Applications []Application
