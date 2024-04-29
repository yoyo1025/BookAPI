package model

import "time"

type Picture struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Image     []byte `json:"image"`	
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

type PictureResponse struct {
	ID uint `json:"id" gorm:"primaryKey"`
}


