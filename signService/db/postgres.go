package db

import (
	"github.com/aliqasemi/ProviderConsumerAttachment/signService/internal/core/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

type DB *gorm.DB
type singleton struct{}

var (
	lock     = &sync.Mutex{}
	database *gorm.DB
	instance *singleton
)

func ConnectPostgres(dsn string) error {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &singleton{}
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("can not connect to database")
		}
		database = db
	}
	return nil
}

func GetDataBase() DB {
	return database
}

func MigratePostgres() {
	entitiesToMigrate := []any{
		&entities.User{},
	}
	if migrate := database.AutoMigrate(entitiesToMigrate...); migrate != nil {
		panic("can not migrate the entities")
	}
}
