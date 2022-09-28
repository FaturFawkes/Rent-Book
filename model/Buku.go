package model

import (
	"gorm.io/gorm"
)

type Buku struct {
	gorm.Model
	ID_user    uint   `gorm:"type:int(11)"`
	Judul      string `gorm:"type:varchar(50)"` 
	Penulis    string `gorm:"type:varchar(255)"`
	Penerbit   string `gorm:"type:varchar(50)"`
	Th_terbit  string `gorm:"type:char(4)"`
	Rents []Rent `gorm:"foreignKey:Id_buku"`
}

type BukuModel struct {
	 DB *gorm.DB
}

func (bm BukuModel) GetAll() ([]Buku, error) {
	var res []Buku
	err := bm.DB.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (bm BukuModel) TambahBuku(book Buku) (Buku, error) {

	err := bm.DB.Create(&book).Error
	if err != nil {
		return Buku{}, err
	}
	return book, nil
}

func (bm BukuModel) UpdateBuku(book Buku) (Buku, error) {
	err := bm.DB.Save(&book).Error
	if err != nil {
		return Buku{}, err
	}
	return book, nil
}

func (bm BukuModel) MyBook(id uint) ([]Buku, error) {
	// var book Buku
	// fmt.Println(book.ID_user)
	var res []Buku
	err := bm.DB.Where("id_user = ?", id).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (bm BukuModel) DeleteBuku(id uint) (Buku, error) {
	err := bm.DB.Delete(&Buku{}, id).Error
	if err != nil {
		return Buku{}, err
	}
	return Buku{}, nil
}