package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func conn() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/rent-book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func main() {
	_, err := conn()
	if err != nil {
		fmt.Println("error")
	} else { 
		fmt.Println("sukses")
	}
}