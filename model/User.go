package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string `gorm:"type:varchar(255)"`
	Username string `gorm:"type:varchar(20)"`
	Password string `gorm:"type:varchar(20)"`
	Email    string `gorm:"type:varchar(50)"`
	Alamat   string `gorm:"type:varchar(255)"`
	Bukus    []Buku `gorm:"foreignKey:ID_user"`
	Rents    []Rent `gorm:"foreignKey:Id_user"`
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

func (us UserModel) Insert(newRegister User) (User, error) {
	err := us.DB.Create(&newRegister).Error
	if err != nil {
		fmt.Println("error on insert", err.Error())
		return User{}, err
	}
	return newRegister, nil
}

func (us UserModel) Update(updateUser User) (User, error) {
	err := us.DB.Save(&updateUser).Error
	if err != nil {
		return User{}, err
	}
	return updateUser, nil
}

func (us UserModel) Delete(deleteUser User) (User, error) {
	err := us.DB.Delete(&deleteUser).Error
	if err != nil {
		return User{}, err
	}
	return deleteUser, nil
}
