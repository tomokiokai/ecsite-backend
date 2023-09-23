package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IShopValidator interface {
	ShopValidate(shop model.Shop) error
}

type shopValidator struct{}

func NewShopValidator() IShopValidator {
	return &shopValidator{}
}

func (sv *shopValidator) ShopValidate(shop model.Shop) error {
	return validation.ValidateStruct(&shop,
		validation.Field(
			&shop.Name,
			validation.Required.Error("name is required"),
			validation.RuneLength(1, 100).Error("limited max 100 char"),
		),
		validation.Field(
			&shop.Address,
			validation.Required.Error("address is required"),
			validation.RuneLength(1, 255).Error("limited max 255 char"),
		),
		validation.Field(
			&shop.Area,
			validation.Required.Error("area is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
		validation.Field(
			&shop.Genre,
			validation.Required.Error("genre is required"),
			validation.RuneLength(1, 50).Error("limited max 50 char"),
		),
		validation.Field(
			&shop.Description,
			validation.Required.Error("description is required"),
			validation.RuneLength(1, 500).Error("limited max 500 char"),
		),
	)
}
