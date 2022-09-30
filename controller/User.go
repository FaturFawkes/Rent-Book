package controller

import (
	"rentbook/model"
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

func (us UserController) Insert(data model.User) (model.User, error) {
	res, err := us.User.Insert(data)
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}

func (us UserController) Update(data model.User) (model.User, error) {
	res, err := us.User.Update(data)
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}

func (us UserController) Delete(data model.User) (model.User, error) {
	res, err := us.User.Delete(data)
	if err != nil {
		return model.User{}, err
	}
	return res, nil
}
