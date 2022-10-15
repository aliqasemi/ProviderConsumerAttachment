package ports

import "github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/core/entities"

type UserInterface interface {
	Index() ([]entities.User, error)
	Create(entities.User) (entities.User, error)
}
