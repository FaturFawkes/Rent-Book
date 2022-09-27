package model

import (
	"gorm.io/gorm"
)
type User struct {
	gorm.Model
	Nama       string `gorm:"type:varchar(255)"`
	Username   string `gorm:"type:varchar(20)"`
	Password   string `gorm:"type:varchar(20)"`
	Email      string `gorm:"type:varchar(50)"`
	Alamat     string `gorm:"type:varchar(255)"`
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
