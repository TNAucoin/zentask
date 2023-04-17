package db

import (
	"fmt"
	"log"

	"github.com/tnaucoin/zentask/authentication-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbUser  = "pg"
	dbPass  = "password"
	dbTable = "users"
	dbPort  = "5432"
)

func Init() *gorm.DB {
	dbURL := fmt.Sprintf("host=db port=%s user=%s password=%s dbname=users sslmode=disable timezone=UTC connect_timeout=5", dbPort, dbUser, dbPass)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
