package controller

import (
	"rentbook/model"
)

type RentController struct {
	Model model.RentModel
}

func (rc RentController) GetAll() ([]model.Rent, error) {

}

func (rc RentController) AddRent() (model.Rent, error) {
	
}