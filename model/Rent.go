package model

import (
	"time"
	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	Id_user     int 
	Id_buku     int 
	Tgl_pinjam  time.Time
	Tgl_kembali time.Time
}

type RentModel struct {
	DB *gorm.DB
}