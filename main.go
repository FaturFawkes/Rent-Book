package main

import (
	"fmt"
	"os"
	"os/exec"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func conn() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/rent-book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
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

func main() {
	var isRunning bool = true
	var inputBook int
	_, err := conn()
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("sukses")
	}

	if isRunning {
		fmt.Println("\t ---- Rent Book ----")
		fmt.Println("Login")
		fmt.Scanln(Login())
		fmt.Println("Login Berhasil => ENTER")
		fmt.Scanln(&inputBook)
		fmt.Println("2. Search Buku")
		fmt.Println("3. Keseluruhan Buku")
		fmt.Println("4. Tambah Buku")
		fmt.Println("5. Pinjam Buku")
		fmt.Println("6. Pencarian Buku")
		fmt.Println("7. Detail Pinjam Buku")
		fmt.Println("8. Pengembalian Buku")
		fmt.Println("9. Logout")
		fmt.Println("Masukkan Input")
		fmt.Scanln(&inputBook)
		switch inputBook {
		case 2:
		case 3:
		case 4:
		case 5:
		case 6:
		case 7:
		case 8:
		case 9:
			Logout()
			isRunning = false
			fmt.Println("\t---Sampai Berjumpa Kembali---\n")
			fmt.Println("\t ---- Rent Book ----")
			fmt.Println("Login")
			fmt.Scanln(Login())
			fmt.Println("Login Berhasil => ENTER")
			fmt.Scanln(&inputBook)
			fmt.Println("2. Search Buku")
			fmt.Println("3. Keseluruhan Buku")
			fmt.Println("4. Tambah Buku")
			fmt.Println("5. Pinjam Buku")
			fmt.Println("6. Pencarian Buku")
			fmt.Println("7. Detail Pinjam Buku")
			fmt.Println("8. Pengembalian Buku")
			fmt.Println("9. Logout")
			fmt.Println("Masukkan Input")
			fmt.Scanln(&inputBook)
		}
	}
}
