package model

import "time"

type Book struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" gorm:"not null"`
	Comment   string `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User User `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"` 
	UserId uint `json:"user_id" gorm:"not null"`
}

type BookResponse struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"not null"`
	Comment   string `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Pictures  []Picture   `json:"pictures"` 
}