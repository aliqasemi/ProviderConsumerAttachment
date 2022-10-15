package repositories

import (
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/db"
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/core/entities"
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/core/ports"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func UserRepositoryBuilder() ports.UserInterface {
	return UserRepository{db: db.GetDataBase()}
}

func (repo UserRepository) Create(user entities.User) (entities.User, error) {
	result := repo.db.Create(&user)
	return user, result.Error
}

func (repo UserRepository) Index() (user []entities.User, err error) {
	result := repo.db.Take(&user)
	return user, result.Error
}
