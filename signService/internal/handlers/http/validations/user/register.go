package validation

import (
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/core/entities"
	"github.com/go-playground/validator"
)

type (
	UserInput struct {
		FirstName       string `json:"first-name" form:"first-name" query:"first-name" param:"first-name" bson:"first-name,omitempty" validate:"required"`
		LastName        string `json:"last-name" form:"last-name" query:"last-name" param:"last-name" bson:"last-name,omitempty" validate:"required"`
		Password        string `json:"password" form:"password" query:"password" param:"password" bson:"password,omitempty" validate:"min=6,eqfield=ConfirmPassword"`
		ConfirmPassword string `json:"confirm-password" form:"confirm-password" query:"confirm-password" param:"confirm-password" bson:"confirm-password,omitempty"`
		PhoneNumber     uint   `json:"phone-number" form:"phone-number" query:"phone-number" param:"phone-number" bson:"phone-number,omitempty" validate:"min=10"`
		Role            string `json:"role" form:"role" query:"role" param:"role" bson:"role,omitempty" validate:""`
		Email           string `json:"email" form:"email" query:"email" param:"email" bson:"email,omitempty" validate:"required,email"`
	}
	UserValidation interface {
		validate(input *UserInput) (bool, error)
		buildEntity(input *UserInput) entities.User
	}
	RegisterValidator struct {
		validator *validator.Validate
	}
)

func (input *UserInput) ValidateAndBuildEntity() (entities.User, error) {
	validatorInput := &RegisterValidator{validator: validator.New()}
	validate, err := validatorInput.validate(input)
	if validate {
		return validatorInput.buildEntity(input), nil
	} else {
		return entities.User{}, err.(validator.ValidationErrors)
	}
}

func (validator *RegisterValidator) validate(input *UserInput) (bool, error) {
	if err := validator.validator.Struct(input); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (validator *RegisterValidator) buildEntity(input *UserInput) entities.User {
	entity := entities.User{
		Firstname:   input.FirstName,
		Lastname:    input.LastName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
	}
	if input.Role == "" {
		entity.Role = "user"
	} else {
		entity.Role = input.Role
	}
	return entity
}
