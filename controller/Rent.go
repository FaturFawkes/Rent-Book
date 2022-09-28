package controller

import (
	"rentbook/model"
)

type RentController struct {
	Model model.RentModel
}

// func (rc RentController) GetAll() ([]model.Rent, error) {

// }

func (rc RentController) AddRent(bookId, userId uint) (model.Result, error) {
	res, err := rc.Model.AddRent(bookId, userId )
	if err != nil {
		return model.Result{}, err
	}
	return res, nil
}

func (rc RentController) CekRent(bookId uint) bool {
	res := rc.Model.CekRent(bookId)
	return res
	// if err != nil {
	// 	return nil, err
	// }
	// return res, nil
}