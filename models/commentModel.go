package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Name    string
	Message string
	PostId  string
}
