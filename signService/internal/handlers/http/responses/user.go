package responses

import (
	"fmt"
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/core/entities"
)

type UserResponse struct {
	ID          uint
	Name        string
	Email       string
	PhoneNumber uint
	Role        string
}

func User(user entities.User) UserResponse {
	return UserResponse{
		ID:          user.ID,
		Name:        fmt.Sprintf("%s %s", user.Firstname, user.Lastname),
		Email:       user.Email,
		PhoneNumber: user.ID,
		Role:        user.Role,
	}
}

func UserCollection(users []entities.User) []UserResponse {
	var responses []UserResponse
	for _, user := range users {
		responses = append(responses, User(user))
	}
	return responses
}
