package gorm

import (
	"github.com/hallex-abreu/golang-clean-architecture/api/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "docker"
const DB_PASSWORD = "docker"
const DB_NAME = "customerDb"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var DB *gorm.DB

func Connection() {
	dns := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&domain.Customer{})

	DB = db
}
