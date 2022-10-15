package ports

import "github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/core/entities"

type userInterface interface {
	Index() ([]entities.User, error)
	Show(id string) (entities.User, error)
	Create(user entities.User) error
	Update(user entities.User, id string) error
	Delete(id string) error
	FindUserByPhoneForLogin(email string) (entities.User, error)
}
