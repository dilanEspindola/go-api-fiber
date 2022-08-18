package database

import (
	"log"
	"os"

	"github.com/dilanEspindola/restapiFiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbInstace struct {
	Db *gorm.DB
}

var Database DbInstace

func DbConnection() {
	dsn := "root:dilan@tcp(localhost:3306)/products?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to DB \n", err)
		os.Exit(2)
	}

	log.Println("DB is connected")

	// migrations
	log.Println("Running migrations")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstace{Db: db}
}
