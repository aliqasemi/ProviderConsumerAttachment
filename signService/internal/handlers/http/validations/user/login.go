package validation

import (
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/core/entities"
	"github.com/go-playground/validator"
)

type (
	LoginInput struct {
		PhoneNumber uint   `json:"phone-number" form:"phone-number" query:"phone-number" param:"phone-number" bson:"phone-number,omitempty" validate:"min=10"`
		Password    string `json:"password" form:"password" query:"password" param:"password" bson:"password,omitempty" validate:"required"`
	}
	LoginValidation interface {
		validate(input *UserInput) (bool, error)
		buildEntity(input *UserInput) entities.User
	}
	LoginValidator struct {
		validator *validator.Validate
	}
)

func (input *LoginInput) ValidateAndBuildEntity() (entities.User, error) {
	validatorInput := &LoginValidator{validator: validator.New()}
	validate, err := validatorInput.validate(input)
	if validate {
		return validatorInput.buildEntity(input), nil
	} else {
		return entities.User{}, err.(validator.ValidationErrors)
	}
}

func (validator *LoginValidator) validate(input *LoginInput) (bool, error) {
	if err := validator.validator.Struct(input); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (validator *LoginValidator) buildEntity(input *LoginInput) entities.User {
	entity := entities.User{
		PhoneNumber: input.PhoneNumber,
	}
	return entity
}
