package controller

import (
	"rentbook/model"
)

type RentController struct {
	Model model.RentModel
}

func (rc RentController) GetAll() ([]model.Rent, error) {
	res, err := rc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// func (rc RentController) AddRent() (model.Rent, error) {

// }
