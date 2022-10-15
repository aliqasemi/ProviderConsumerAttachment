package responses

import "github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/core/entities"

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
		Name:        user.Firstname + user.Lastname,
		Email:       user.Email,
		PhoneNumber: user.ID,
		Role:        user.Role,
	}
}
