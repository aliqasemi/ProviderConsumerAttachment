package ports

import "github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/core/entities"

type UserInterface interface {
	Index() ([]entities.User, error)
	Create(entities.User) (entities.User, error)
	Find(uint) (entities.User, error)
	Show(uint) (entities.User, error)
	Update(uint, entities.User) (entities.User, error)
	Delete(uint) error
}
