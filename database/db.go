package database

import (
	"fmt"

	"final-project/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	POSTGRES_HOST = "localhost"
	POSTGRES_PORT = 5432
	POSTGRES_DB   = "final-project-db"
	POSTGRES_USER = "postgres"
	POSTGRES_PASS = "postgres"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		POSTGRES_HOST, POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASS, POSTGRES_DB,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.Social{})

	return db
}
