package controller

import (
	"fmt"
	"rentbook/model"
)

type RentController struct {
	Model model.RentModel
}

func (rc RentController) AddRent(bookId, userId uint) (model.Result, error) {
	res, err := rc.Model.AddRent(bookId, userId )
	if err != nil {
		return model.Result{}, err
  }
  return res, nil
}

func (rc RentController) GetAll() ([]model.Rent, error) {
	res, err := rc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (rc RentController) CekRent(bookId uint) bool{
	rent, id, err := rc.Model.CekRent(bookId)
	if err != nil {
		fmt.Println("error rent ln 27")
	}

	if rent != nil {
		if rent[0].Tgl_pinjam == rent[0].Tgl_kembali {
			return false
		}
	} else if id != nil {
		return true
	}

	return true
}
