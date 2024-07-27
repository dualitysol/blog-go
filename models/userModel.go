package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
	Age       int
	Gender    string
	Address   string
	Website   string
}

// func (u *User) ComparePassword(p *string) bool {
// 	return true
// }
