package validation

import (
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/core/entities"
	"github.com/go-playground/validator"
)

type (
	UserUpdateInput struct {
		FirstName   string `json:"first-name" form:"first-name" query:"first-name" param:"first-name" bson:"first-name,omitempty" validate:"required"`
		LastName    string `json:"last-name" form:"last-name" query:"last-name" param:"last-name" bson:"last-name,omitempty" validate:"required"`
		PhoneNumber uint   `json:"phone-number" form:"phone-number" query:"phone-number" param:"phone-number" bson:"phone-number,omitempty" validate:"min=10"`
		Role        string `json:"role" form:"role" query:"role" param:"role" bson:"role,omitempty" validate:""`
		Email       string `json:"email" form:"email" query:"email" param:"email" bson:"email,omitempty" validate:"required,email"`
	}
	UserUpdateValidation interface {
		validate(input *UserUpdateInput) (bool, error)
		buildEntity(input *UserUpdateInput) entities.User
	}
	UserUpdateValidator struct {
		validator *validator.Validate
	}
)

func (input *UserUpdateInput) ValidateAndBuildEntity() (entities.User, error) {
	validatorInput := &UserUpdateValidator{validator: validator.New()}
	validate, err := validatorInput.validate(input)
	if validate {
		return validatorInput.buildEntity(input), nil
	} else {
		return entities.User{}, err.(validator.ValidationErrors)
	}
}

func (validator *UserUpdateValidator) validate(input *UserUpdateInput) (bool, error) {
	if err := validator.validator.Struct(input); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (validator *UserUpdateValidator) buildEntity(input *UserUpdateInput) entities.User {
	entity := entities.User{
		Firstname:   input.FirstName,
		Lastname:    input.LastName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Role:        input.Role,
	}
	return entity
}
