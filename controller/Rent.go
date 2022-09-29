package controller

import (
	"fmt"
	"rentbook/model"
)

type RentController struct {
	Model model.RentModel
}

func (rc RentController) AddRent(bookId, userId uint) (model.Rent, error) {
	res, err := rc.Model.AddRent(bookId, userId)
	if err != nil {
		return model.Rent{}, err
	}
	return res, nil
}

func (rc RentController) GetAll() ([]model.DetailRent, error) {
	// Query get all table rent
	res, err := rc.Model.GetAll()
	if err != nil {
		fmt.Println("error get all rent")
	}
  return  res, nil

	// CEK BUKU SUDAH DIKEMBALIKAN ?
}


// func (rc RentController) CekRent(bookId uint) bool{
// var finalRes []model.Rent
// 	for i := 0; i < len(res); i++ {
// 		kembali := res[i]
// 		if kembali.Tgl_kembali == kembali.Tgl_pinjam {
// 			finalRes = append(finalRes, res[i])
// 		}
// 	}

// 	return finalRes, nil
// }

func (rc RentController) CekRent(bookId uint) bool {
	rent, id, err := rc.Model.CekRent(bookId)
	if err != nil {
		fmt.Println("error rent ln 27")
	}

	if rent != nil {
		if rent[0].Status == "" {
			return false
		} 
		if rent[0].Id_buku == bookId{
			return false
		}
	} else if id != nil {
		return true
	}

	return true
}

func (rc RentController)  KembaliBuku(idBuku, idUser uint) (bool, error){
	res, err := rc.Model.KembaliBuku(idBuku, idUser)
	if err != nil {
		return res, err
	}
	return res, nil
}

// func (rc RentController) AllBookRent()([]model.Buku, error) {
// 	res, err := rc.Model.AllBookRent()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

// func (rc RentController) UpdateTgl(id uint) ([]model.Rent, error) {
// 	_, err := rc.Model.UpdateTgl(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return []model.Rent{}, nil
// }

