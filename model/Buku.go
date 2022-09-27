package model

import (
	"gorm.io/gorm"
)

type Buku struct {
	gorm.Model
	ID_user    uint  
	Judul      string `gorm:"type:varchar(50)"` 
	Penulis    string `gorm:"type:varchar(255)"`
	Penerbit   string `gorm:"type:varchar(50)"`
	Th_terbit  string `gorm:"type:char(4)"`
}

type BukuModel struct {
	 DB *gorm.DB
}