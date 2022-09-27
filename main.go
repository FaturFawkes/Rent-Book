package main

import (
	"fmt"
	"os"
	"os/exec"
	"rentbook/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Buku{})
	// db.AutoMigrate(&model.Rent{})
	db.AutoMigrate(&model.User{})
}

func conn() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/rent-book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Logout() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Login() (username string, password string) {
	var us, pw string
	fmt.Println("Masukkan User Name : ")
	fmt.Scanln(&us)
	fmt.Println("Masukkan Password : ")
	fmt.Scanln(&pw)

	return us, pw
}

func Register() (nama, email, password string) {
	var na, em, pas string
	fmt.Println("Masukkan User Name : ")
	fmt.Scanln(&na)
	fmt.Println("Masukkan Email : ")
	fmt.Scanln(&em)
	fmt.Println("Masukkan Password : ")
	fmt.Scanln(&pas)

	return na, em, pas
}

func main() {
	var isRunning bool = true
	var inputBook int
	var inputBook2 int
	_, err := conn()
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("sukses")
	}

	for isRunning {
		fmt.Println("\t ---- Rent Book ----")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Search Buku")
		fmt.Println("4. Keseluruhan Buku\n")
		fmt.Println("\tMasukkan Piliha Anda")
		fmt.Scanln(&inputBook)
		switch inputBook {
		case 1:
			fmt.Println("Silahkan Login\n")
			Login()
			fmt.Println("\tLogin Berhasil\n")
		case 2:
			fmt.Println("Silahkan Registrasi\n")
			Register()
			fmt.Println("\tRegister Berhasil\n")
		}
		fmt.Println("3. Search Buku")
		fmt.Println("4. Keseluruhan Buku")
		fmt.Println("5. Tambah Buku")
		fmt.Println("6. Pinjam Buku")
		fmt.Println("7. Pencarian Buku")
		fmt.Println("8. Detail Pinjam Buku")
		fmt.Println("9. Pengembalian Buku")
		fmt.Println("10. Logout")
		fmt.Println("Masukkan Input")
		fmt.Scanln(&inputBook)
		Logout()
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Search Buku")
		fmt.Println("4. Keseluruhan Buku\n")
		fmt.Println("\tMasukkan Piliha Anda")
		Login()
		fmt.Scanln(&inputBook2)
		fmt.Println("3. Search Buku")
		fmt.Println("4. Keseluruhan Buku")
		fmt.Println("5. Tambah Buku")
		fmt.Println("6. Pinjam Buku")
		fmt.Println("7. Pencarian Buku")
		fmt.Println("8. Detail Pinjam Buku")
		fmt.Println("9. Pengembalian Buku")
		fmt.Println("10. Logout")
		fmt.Println("Masukkan Input")
		fmt.Scanln(&inputBook2)
		switch inputBook2 {
		case 1:
			fmt.Println("Silahkan Login\n")
			Login()
			fmt.Println("\tLogin Berhasil\n")
		case 2:
			fmt.Println("Silahkan Registrasi\n")
			Register()
			fmt.Println("\tRegister Berhasil\n")
		case 3:
		case 4:
		case 5:
		case 6:
		case 7:
		case 8:
		case 9:
		case 10:
			Logout()
			// fmt.Println("Login Berhasil => ENTER")
			// fmt.Scanln(&inputBook2)
			// fmt.Println("\t ---- Rent Book ----")
			// fmt.Println("1. Login")
			// fmt.Println("2. Register")
			// fmt.Println("3. Search Buku")
			// fmt.Println("4. Keseluruhan Buku\n")
			// fmt.Println("\tMasukkan Piliha Anda")
			// fmt.Scanln(&inputBook2)
			// fmt.Println("3. Search Buku")
			// fmt.Println("4. Keseluruhan Buku")
			// fmt.Println("5. Tambah Buku")
			// fmt.Println("6. Pinjam Buku")
			// fmt.Println("7. Pencarian Buku")
			// fmt.Println("8. Detail Pinjam Buku")
			// fmt.Println("9. Pengembalian Buku")
			// fmt.Println("10. Logout")
			// fmt.Println("Masukkan Input")
			// fmt.Scanln(&inputBook2)
		}
	}
}
