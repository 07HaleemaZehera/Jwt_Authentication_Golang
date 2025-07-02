package database

import (
	"JWT_AUTHENTICATION/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDB() (*gorm.DB, error) {
dsn := "root:Newpassword2025@tcp(localhost:3306)/auth?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err 
	}
	DB = db
	 db.AutoMigrate(&models.User{})

	 return db , nil
}