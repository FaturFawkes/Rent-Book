package controller

import (
	"rentbook/model"
	// "golang.org/x/text/date"
)

type UserController struct {
	User model.UserModel
}

func (us UserController) Login(username, password string) ([]model.User, error) {
	res, err := us.User.Login(username, password)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// main(panggil func controller) ->func controller panggil func model -> func model query data

func (us UserController) Insert(data model.User) (model.User, error) {
	res, err := us.User.Insert(data)
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}
