package entity

import "gorm.io/gorm"

type Commands struct {
	gorm.Model
	Cmd  string
	Desc string
}
