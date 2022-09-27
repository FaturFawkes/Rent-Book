package controller

import "rentbook/model"

type BukuController struct {
	Model model.BukuModel
}

func (bc BukuController) GetAll() ([]model.Buku, error){
	res, err := bc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (bc BukuController) TambahBuku(buku model.Buku) (model.Buku, error) {
	res, err := bc.Model.TambahBuku(buku) 
	if err != nil {
		return model.Buku{}, err
	}
	return res, nil
}