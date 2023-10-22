package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user model.User) error
	UserLoginValidate(user model.User) error  // 新しいバリデーション関数のインターフェース
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		// emailのバリデーション
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
			is.Email.Error("is not valid email format"),
		),
		// passwordのバリデーション
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("limited min 6 max 30 char"),
		),
		// nameのバリデーション（新規追加）
		validation.Field(
			&user.Name,
			validation.Required.Error("name is required"), // 名前が必須であることを示す
			validation.RuneLength(1, 50).Error("name must be between 1 and 50 characters"), // 文字数の制限
		),
	)
}

// UserLoginValidateはログイン時のバリデーションを行います。ここではnameは検証しません。
func (uv *userValidator) UserLoginValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		// Emailのバリデーション
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("invalid email format"),
		),
		// Passwordのバリデーション
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 100).Error("password must be between 6 and 100 characters"),
		),
		// Nameのバリデーションはここでは行いません。
	)
}

