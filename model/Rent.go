package model

import (
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
  	Id_user     uint `gorm:"type:int(11)"`
	Id_buku     uint `gorm:"type:int(11)"` 
	Tgl_pinjam  time.Time `gorm:"autoCreateTime"`
	Tgl_kembali time.Time `gorm:"autoUpdateTime"`
}

type RentModel struct {
	DB *gorm.DB
}

func (rm RentModel) GetAll() ([]Rent, error) {
	var res []Rent
	err := rm.DB.Session(&gorm.Session{QueryFields: true}).Model(&Rent{}).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (rm RentModel) AddRent(bookId, userId uint) (Result, error){

	// err := rm.DB.Model(&Rent{}).Select("users.nama, bukus.judul").Joins("left join users on user.id = ?", userId).Joins("left join bukus on bukus.id = ?", bookId).Scan(&Rent{}).Error
	err := rm.DB.Create(&Rent{Id_user: userId, Id_buku: bookId}).Error
	if err != nil {
		return Result{}, err
	}
	return Result{}, nil
}

func (rm RentModel) CekRent(bookId uint) ([]Rent, []Buku, error){
	var rent []Rent
	var buku []Buku
	
	cek := rm.DB.Where("id_buku = ?", bookId).Find(&rent).RowsAffected
	if cek > 0 {
		return rent, nil, nil
	} else {
		bookCek := rm.DB.Find(&buku, bookId).Error
		if bookCek != nil {
			return nil, nil, bookCek
		}
		return nil, buku, nil
	}
}
