package model

import "time"

type Favorite struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ShopID    uint      `json:"shop_id" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsFavorite bool `json:"is_favorite"`
	Shop   Shop `gorm:"foreignKey:ShopID"`
  User   User `gorm:"foreignKey:UserID"`
}

type FavoriteResponse struct {
	ID        uint      `json:"id"`
	Shop      Shop      `json:"shop"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsFavorite bool `json:"is_favorite"`
}
