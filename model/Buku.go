package model

import (
	"gorm.io/gorm"
)

type Buku struct {
	gorm.Model
	ID_user    uint  
	Judul      string  
	Penulis    string  
	Penerbit   string  
	Th_terbit  string 
}

type BukuModel struct {
	 DB *gorm.DB
}