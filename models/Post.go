package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string    `gorm:"type:varchar(200)"`
	Description string    `gorm:"type:text"`
	Location    string    `gorm:"type:varchar(200)"`
	Date        time.Time `gorm:"type:date"`
}
