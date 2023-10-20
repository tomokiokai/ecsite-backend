package model

import "time"

type Shop struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Address     string    `json:"address" gorm:"not null"`
	Area        string    `json:"area" gorm:"not null"`
	Genre       string    `json:"genre" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Favorites []Favorite `json:"favorites" gorm:"foreignKey:ShopID"`
	Reservations []Reservation `json:"reservations" gorm:"foreignKey:ShopID"`
}

type ShopResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Area        string    `json:"area"`
	Genre       string    `json:"genre"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}