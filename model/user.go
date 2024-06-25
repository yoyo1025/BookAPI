package model

import "time"

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	UserName string `json:"username" gorm:"column:user_name"`
	Email string `json:"email" gorm:"unique"` 
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
	UserName string `json:"username"`
}