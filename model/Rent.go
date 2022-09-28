package model

import (
	"time"
	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	Id_user     int `gorm:"type:int(11)"`
	Id_buku     int `gorm:"type:int(11)"`
	Pemilik     string `gorm:"type:varchar(255)"`
	Judul       string `gorm:"type:varchar(50)"`
	Tgl_pinjam  time.Time
	Tgl_kembali time.Time
}

type RentModel struct {
	DB *gorm.DB
}

func (rm RentModel) GetAll() ([]Rent, error) {

}

func (rm RentModel) AddRent() (Rent, error) {

}