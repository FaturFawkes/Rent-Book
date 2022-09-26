package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama       string 
	Username   string 
	Password   string 
	Email      string 
	Alamat     string
	Status     string
}

type UserModel struct {
	DB *gorm.DB
}
