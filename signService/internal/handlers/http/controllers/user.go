package controllers

import (
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/handlers/http/responses"
	validation "github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/handlers/http/validations/user"
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Index(c echo.Context) error {
	repo := repositories.UserRepositoryBuilder()
	entityList, err := repo.Index()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, responses.UserCollection(entityList))
}

func Register(c echo.Context) error {
	userInput := new(validation.UserInput)
	if err := c.Bind(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	entity, err := userInput.ValidateAndBuildEntity()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err = entity.HashPassword(entity.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	repo := repositories.UserRepositoryBuilder()
	if entity, err = repo.Create(entity); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, responses.User(entity))
}
