package main

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name     string
	Username string
	Password string
}
