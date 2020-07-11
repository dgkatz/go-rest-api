package models

import (
	"time"
)

type Assignment struct {
	ID uint `gorm:"primary_key;"`
	UserID uint
	User User `gorm:"ForeignKey:UserID"`
	Title string
	Description string
	Deadline time.Time
	Status string
}