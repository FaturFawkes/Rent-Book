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

func (us UserModel) Login(username, password string) ([]User, error) {
	var res []User
	err := us.DB.Find(&res, "username = ? AND password = ?", username, password).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
