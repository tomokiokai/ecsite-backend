package repository

import (
	"go-rest-api/model"
	"gorm.io/gorm"
)

type IReservationRepository interface {
    MakeReservation(reservation *model.Reservation) (model.Reservation, error)
    CancelReservation(reservationId string) error
    GetReservationByUser(userId string) ([]model.Reservation, error)
    GetAllReservations() ([]model.Reservation, error)
    UpdateReservation(reservation *model.Reservation) (model.Reservation, error)
    GetReservationsForBuild() ([]model.Reservation, error)
}

type reservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) IReservationRepository {
	return &reservationRepository{db}
}

func (rr *reservationRepository) MakeReservation(reservation *model.Reservation) (model.Reservation, error) {
    result := rr.db.Create(reservation)
    return *reservation, result.Error
}

func (rr *reservationRepository) CancelReservation(reservationId string) error {
    result := rr.db.Delete(&model.Reservation{}, reservationId)
    return result.Error
}

func (rr *reservationRepository) GetReservationByUser(userId string) ([]model.Reservation, error) {
    var reservations []model.Reservation
    result := rr.db.Where("user_id = ?", userId).Find(&reservations)
    return reservations, result.Error
}

func (rr *reservationRepository) GetAllReservations() ([]model.Reservation, error) {
    var reservations []model.Reservation
    result := rr.db.Find(&reservations)
    return reservations, result.Error
}

func (rr *reservationRepository) UpdateReservation(reservation *model.Reservation) (model.Reservation, error) {
    result := rr.db.Save(reservation)
    return *reservation, result.Error
}

func (rr *reservationRepository) GetReservationsForBuild() ([]model.Reservation, error) {
    var reservations []model.Reservation
    result := rr.db.Preload("User").Find(&reservations) // ここでは関連するユーザー情報も取得します
    return reservations, result.Error
}
