package main

import (
	"fmt"
	"os"
	"os/exec"
	"rentbook/controller"
	"rentbook/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Buku{})
	db.AutoMigrate(&model.Rent{})
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
	conn, err := conn()
	if err != nil {
		fmt.Println("error", err.Error())
	}

	migrate(conn)
	session := model.User{}
	userModel := model.UserModel{conn}
	userControll := controller.UserController{userModel}

	for isRunning {
		var inputMenu int
		fmt.Println("")
		fmt.Println("===== RENTAL BUKU =====")
		fmt.Println("")
		fmt.Println("1.Login")
		fmt.Println("2.Update Profile(login)")
		fmt.Println("3.Hapus Profile(login)")
		fmt.Println("4.Daftar Buku")
		fmt.Println("5.Pinjam Buku (login)")
		fmt.Println("6.Pinjam buku teman (login)")
		fmt.Println("7.Register")
		fmt.Println("8.Logout")
		fmt.Println("")
		fmt.Println("Pilih Menu : ")
		fmt.Scanln(&inputMenu)
		switch inputMenu {
		case 1:
			var login model.User
			fmt.Print("Username : ")
			fmt.Scanln(&login.Username)
			fmt.Print("Password :")
			fmt.Scanln(&login.Password)
			res, err := userControll.Login(login.Username, login.Password)
			if err != nil {
				fmt.Println("gagal login")
			} else {
				session = res[0]
				fmt.Println(session.Nama)
			}
		case 2:
			var update model.User
			var cekLogin = true
			if session.ID == 0 {
				cekLogin = false
			}
			if cekLogin == true {
				fmt.Println("Nama : ")
				fmt.Scanln(&update.Nama)
				fmt.Println("Username : ")
				fmt.Scanln(&update.Username)
				fmt.Println("Password : ")
				fmt.Scanln(&update.Password)
				fmt.Println("Email : ")
				fmt.Scanln(&update.Email)
				fmt.Println("Alamat : ")
				fmt.Scanln(&update.Alamat)
				fmt.Println("Status : ")
				fmt.Scanln(&update.Status)
				update.ID = session.ID
				_, err := userControll.Update(update)
				if err != nil {
					fmt.Println("gagal update")
				} else {
					fmt.Println(session.ID)
				}
			}
		case 7:
			var user model.User
			fmt.Println("Nama : ")
			fmt.Scanln(&user.Nama)
			fmt.Println("Username : ")
			fmt.Scanln(&user.Username)
			fmt.Println("Password : ")
			fmt.Scanln(&user.Password)
			fmt.Println("Email : ")
			fmt.Scanln(&user.Email)
			fmt.Println("Alamat : ")
			fmt.Scanln(&user.Alamat)
			fmt.Println("Status : ")
			fmt.Scanln(&user.Status)
			_, err := userControll.Insert(user)
			if err != nil {
				fmt.Println("gagal register")
			} else {
				fmt.Println(session.Nama)
			}
		case 8:
			Logout()
			session = model.User{}

		}
	}
}
