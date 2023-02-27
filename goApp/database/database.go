package database

import (
	"log"
	"os"

	"github.com/Carlfred3010/goApp/models"
	"github.com/glebarez/go-sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstance struct {
	Db *gorm.DB
}

var Database DBinstance

func ConnectDb() {

	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database\n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to database successfully")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running Migrations")
	db.AutoMigrate(&models.User{}, &models.Stock{})

	Database = DBinstance{Db: db}

}
