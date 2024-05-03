package model

import "time"

type Picture struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Image       []byte    `json:"-"` // JSON出力から除外する
	ImageBase64 string    `json:"imageBase64"` // Base64エンコードされた画像データ
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserId      uint      `json:"user_id"`
}

type PictureResponse struct {
	ID uint `json:"id" gorm:"primaryKey"`
	ImageBase64 string    `json:"imageBase64"` // Base64エンコードされた画像データ
}


