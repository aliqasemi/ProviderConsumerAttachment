package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB *gorm.DB

func ConnectPostgres(dsn string) DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("can not connect to database")
	}
	return db
}

func MigratePostgres() {

}
