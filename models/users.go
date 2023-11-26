package models

import "time"

type User struct {
	Id        uint `json:"id" gorm:"primaryKey" `
	CreatedAt time.Time
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname" `
}
