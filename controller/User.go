package controller

import "rentbook/model"

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

