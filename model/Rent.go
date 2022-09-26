package model

import (
	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	Id_user     int 
	Id_buku     int 
	Tgl_pinjam  string
	Tgl_kembali string
}

type RentModel struct {
	DB *gorm.DB
}