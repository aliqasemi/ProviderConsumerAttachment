package ports

import "github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/core/entities"

type UserInterface interface {
	Create(entities.User) (entities.User, error)
}
