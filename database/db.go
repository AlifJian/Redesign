package database

import (
	"backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/gdsc"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Post{})
	return db
}
