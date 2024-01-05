package entity

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	CommandsID uint `gorm:"uniqueIndex"`
}
