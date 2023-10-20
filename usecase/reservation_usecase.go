package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type IReservationUsecase interface {
    MakeReservation(reservation model.Reservation) (model.Reservation, error)
    CancelReservation(reservationId string) error
    GetReservationByUser(userId string) ([]model.Reservation, error)
    GetAllReservations() ([]model.Reservation, error)
    UpdateReservation(reservation model.Reservation) (model.Reservation, error)
    GetReservationsForBuild() ([]model.Reservation, error)
}

type reservationUsecase struct {
    rr repository.IReservationRepository
	rv validator.IReservationValidator // バリデータのインスタンス
}

func NewReservationUsecase(rr repository.IReservationRepository, rv validator.IReservationValidator) IReservationUsecase {
	return &reservationUsecase{rr, rv}
}

func (ru *reservationUsecase) MakeReservation(reservation model.Reservation) (model.Reservation, error) {
    // 入力データのバリデーションを行う
    if err := ru.rv.ReservationValidate(reservation); err != nil {
        return model.Reservation{}, err
    }
    // バリデーションが成功したら、予約を作成
    return ru.rr.MakeReservation(&reservation)
}

func (ru *reservationUsecase) CancelReservation(reservationId string) error {
    return ru.rr.CancelReservation(reservationId)
}

func (ru *reservationUsecase) GetReservationByUser(userId string) ([]model.Reservation, error) {
    return ru.rr.GetReservationByUser(userId)
}

func (ru *reservationUsecase) GetAllReservations() ([]model.Reservation, error) {
    return ru.rr.GetAllReservations()
}

func (ru *reservationUsecase) UpdateReservation(reservation model.Reservation) (model.Reservation, error) {
    // 更新前にもバリデーションを行う
    if err := ru.rv.ReservationValidate(reservation); err != nil {
        return model.Reservation{}, err
    }
    return ru.rr.UpdateReservation(&reservation)
}

func (ru *reservationUsecase) GetReservationsForBuild() ([]model.Reservation, error) {
    return ru.rr.GetReservationsForBuild()
}

