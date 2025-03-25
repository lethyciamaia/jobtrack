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
	ID          uint      `json:"id" gorm:"primaryKey; autoIncrement" `
	Company     string    `json:"company" gorm:"column:company"`
	Description string    `json:"description" gorm:"column:description"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	Status      Status    `json:"status" gorm:"column:status"`
}

type ApplicationUpdate struct {
	ID     uint   `json:"id"`
	Status Status `json:"status"`
}
