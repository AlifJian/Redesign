package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title       string `gorm:"type:varchar(200)"`
	Description string `gorm:"type:text"`
	ImageUrl    string `gorm:"type:varchar(100)"`
	EventUrl    string `gorm:"type:text"`
	Location    string `gorm:"type:varchar(200)"`
	Date        string `gorm:"type:varchar(100)"`
}
