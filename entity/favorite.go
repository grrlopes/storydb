package entity

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	CommandsID uint `gorm:"uniqueIndex"`
}

type FavoriteView struct {
	CommandsID   int    `gorm:"column:commands_id"`
	CommandsCMD  string `gorm:"column:cmd"`
	CommandsDESC string `gorm:"column:desc"`
	FavoriteID   string `gorm:"column:favorite_id"`
}
