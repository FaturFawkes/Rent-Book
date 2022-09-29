package model

import (
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	Id_user     uint      `gorm:"type:int(11)"`
	Id_buku     uint      `gorm:"type:int(11)"`
	Tgl_pinjam  time.Time `gorm:"autoCreateTime"`
	Tgl_kembali time.Time `gorm:"autoUpdateTime"`
	Status string `gorm:"type:varchar(20)"`
}

type DetailRent struct {
	Id uint
	Judul string
	Nama string
}

type RentModel struct {
	DB *gorm.DB
}

func (rm RentModel) GetAll() ([]DetailRent, error) {
	var res []DetailRent

	// err1 := rm.DB.Select("bukus.id, bukus.judul, users.nama").Table("bukus").
	// 	Joins("left join users on bukus.id_user = users.id").
	// 	Joins("Left join rents on rents.id_buku = bukus.id").
	// 	Where("rents.status = ?", "kembali").Scan(&res).Error
		
	err2 := rm.DB.Select("bukus.id, bukus.judul, users.nama").Table("bukus").
			Joins("left join rents on bukus.id != rents.id_bukus"). 
			Error
	// if err1 != nil {
	// 	return nil, err1
	// }
	if err2 != nil {
		return nil, err2
	}


	return res, err2

}

func (rm RentModel) AddRent(bookId, userId uint) (Rent, error) {

	// err := rm.DB.Model(&Rent{}).Select("users.nama, bukus.judul").Joins("left join users on user.id = ?", userId).Joins("left join bukus on bukus.id = ?", bookId).Scan(&Rent{}).Error
	err := rm.DB.Create(&Rent{Id_user: userId, Id_buku: bookId}).Error
	if err != nil {
		return Rent{}, err
	}
	return Rent{}, nil
}

func (rm RentModel) CekRent(bookId uint) ([]Rent, []Buku, error) {
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
func (rm RentModel) KembaliBuku(idBuku, idUser uint) (bool, error){
	var res Rent
	err := rm.DB.Model(res).Where("id_buku = ?", idBuku).Where("id_user = ?", idUser).Updates(Rent{Status: "kembali", Tgl_kembali: time.Now()}).Error
	err1 := rm.DB.Where("id_buku = ?", idBuku).Where("id_user = ?", idUser).Delete(&res).Error
	if err != nil {
		return false, err
	}
	if err1 != nil {
		return false, err
	}
	return true, nil
}

// func (rm RentModel) AllBookRent() ([]Buku, error) {
// 	var res []Buku
// 	err := rm.DB.
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

func (rm RentModel) UpdateTgl(id uint) (int, error) {
	var time = time.Now()
	err := rm.DB.Model(&Rent{}).Where("id_buku = ?", id).Update("tgl_kembali", time).RowsAffected
	if err > 0 {
		return 1, nil
	}
	return 0, nil
}
