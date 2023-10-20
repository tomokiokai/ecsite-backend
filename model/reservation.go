package model

import "time"

type Reservation struct {
    ID      uint      `json:"id" gorm:"primaryKey"`
    Date    time.Time `json:"date" gorm:"not null"`
    Time    string    `json:"time" gorm:"not null"`
    ShopID  uint      `json:"shop_id" gorm:"not null"`
    UserID  uint      `json:"user_id" gorm:"not null"`
    Num     int       `json:"num" gorm:"not null"`
}

type ReservationResponse struct {
    ID     uint      `json:"id"`
    Date   time.Time `json:"date"`
    Time   string    `json:"time"`
    ShopID uint      `json:"shop_id"`
    UserID uint      `json:"user_id"`
    Num    int       `json:"num"`
}
