package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IBlogValidator interface {
	BlogValidate(blog model.Blog) error
}

type blogValidator struct{}

func NewBlogValidator() IBlogValidator {
	return &blogValidator{}
}

func (bv *blogValidator) BlogValidate(blog model.Blog) error {
	return validation.ValidateStruct(&blog,
		validation.Field(
			&blog.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 100).Error("limited max 100 char"), // Assuming a longer title for blogs
		),
		validation.Field(
			&blog.Content,
			validation.Required.Error("content is required"),
			validation.RuneLength(1, 5000).Error("limited max 5000 char"), // Assuming a longer content for blogs
		),
	)
}
