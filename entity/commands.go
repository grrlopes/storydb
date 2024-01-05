package entity

import "gorm.io/gorm"

type Commands struct {
	gorm.Model
	Cmd      string `gorm:"uniqueIndex"`
	Desc     string
	Favorite Favorite
}
