package models

import (
	"time"
)

type Assignment struct {
	ID uint `gorm:"primary_key;"`
	User User
	Title string
	Description string
	Deadline time.Time
	Status string
}