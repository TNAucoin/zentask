package db

import (
	"log"
	"os"

	"github.com/tnaucoin/zentask/authentication-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := os.Getenv("DSN")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
