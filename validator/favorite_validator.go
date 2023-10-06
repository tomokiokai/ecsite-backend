package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IFavoriteValidator interface {
	FavoriteValidate(favorite model.Favorite) error
}

type favoriteValidator struct{}

func NewFavoriteValidator() IFavoriteValidator {
	return &favoriteValidator{}
}

func (fv *favoriteValidator) FavoriteValidate(favorite model.Favorite) error {
	return validation.ValidateStruct(&favorite,
		validation.Field(
			&favorite.UserID,
			validation.Required.Error("user ID is required"),
		),
		validation.Field(
			&favorite.ShopID,
			validation.Required.Error("shop ID is required"),
		),
	)
}