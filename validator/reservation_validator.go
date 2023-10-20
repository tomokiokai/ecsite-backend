package validator

import (
	"go-rest-api/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IReservationValidator interface {
	ReservationValidate(reservation model.Reservation) error
}

type reservationValidator struct{}

func NewReservationValidator() IReservationValidator {
	return &reservationValidator{}
}

func (rv *reservationValidator) ReservationValidate(reservation model.Reservation) error {
	return validation.ValidateStruct(&reservation,
		// UserIDは必須
		validation.Field(&reservation.UserID, validation.Required.Error("user ID is required")),
		// ShopIDも必須
		validation.Field(&reservation.ShopID, validation.Required.Error("shop ID is required")),
		// Dateは必須
		validation.Field(&reservation.Date, validation.Required.Error("date is required")),
		// Timeは必須
		validation.Field(&reservation.Time, validation.Required.Error("time is required")),
		// Num (予約人数) も必須
		validation.Field(&reservation.Num, validation.Required.Error("number of people is required")),
		// ここで他のバリデーションルールを追加できます。例えば、予約日が未来であることを確認するなど。
	)
}

