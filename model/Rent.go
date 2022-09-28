package model

import (
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	Id_user     int    `gorm:"type:int(11)"`
	Id_buku     int    `gorm:"type:int(11)"`
	Pemilik     string `gorm:"type:varchar(255)"`
	Judul       string `gorm:"type:varchar(50)"`
	Tgl_pinjam  time.Time
	Tgl_kembali time.Time
}

type RentModel struct {
	DB *gorm.DB
}

// db.Session(&gorm.Session{QueryFields: true}).Find(&user)
func (rm RentModel) GetAll() ([]Rent, error) {
	var res []Rent
	err := rm.DB.Session(&gorm.Session{QueryFields: true}).Model(&Rent{}).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Select("id_user", "id_buku", "pemilik", "judul").Model(&Rent{}).Find(&res).Error
// func (rm RentModel) AddRent() (Rent, error) {

// }
