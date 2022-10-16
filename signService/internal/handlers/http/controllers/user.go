package controllers

import (
	"errors"
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/core/services"
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/handlers/http/responses"
	validation "github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/handlers/http/validations/user"
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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

func Login(c echo.Context) error {
	loginInput := new(validation.LoginInput)
	if err := c.Bind(loginInput); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	entity, err := loginInput.ValidateAndBuildEntity()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	repo := repositories.UserRepositoryBuilder()
	if entity, err = repo.Find(entity.PhoneNumber); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err = entity.CheckPassword(loginInput.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, errors.New("invalid password"))
	}

	token, err := services.GenerateJWT(entity.Email, strconv.Itoa(int(entity.PhoneNumber)), strconv.Itoa(int(entity.ID)), entity.Role)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
		"user":  responses.User(entity),
	})
}

func Show(c echo.Context) error {
	repo := repositories.UserRepositoryBuilder()
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	entity, err := repo.Show(uint(id))
	return c.JSON(http.StatusOK, responses.User(entity))
}

func Auth(c echo.Context) error {
	authContext := c.(*services.AuthContext)
	repo := repositories.UserRepositoryBuilder()
	id, err := strconv.ParseUint(authContext.GetUserId(), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	entity, err := repo.Show(uint(id))
	return c.JSON(http.StatusOK, responses.User(entity))
}
