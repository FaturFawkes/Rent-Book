package main

import (
	"fmt"
	"os"
	"os/exec"
	"rentbook/controller"
	"rentbook/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	// db.AutoMigrate(&model.Buku{})
	// db.AutoMigrate(&model.Rent{})
	// db.AutoMigrate(&model.User{})

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

func clear(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clear()
	var isRunning bool = true
	conn, err := conn()
	if err != nil {
		fmt.Println("error", err.Error())
	}
	
	// migrate(conn)
	session := model.User{}
	userModel := model.UserModel{conn}
	userControll := controller.UserController{userModel}
	bukuModel := model.BukuModel{conn}
	bukuControll := controller.BukuController{bukuModel}

	for isRunning {
		var inputMenu int
		fmt.Println("")
		fmt.Println("===== RENTAL BUKU =====")
		fmt.Println("")
		fmt.Println("1.Login")
		fmt.Println("2.Update Profile(login)")
		fmt.Println("3.Buku Saya (login)")
		fmt.Println("4.Daftar Buku")
		fmt.Println("5.Pinjam Buku (login)")
		fmt.Println("6.Lihat Buku Saya (login)")
		fmt.Println("7.Register")
		fmt.Println("8.Logout")
		fmt.Println("")
		fmt.Print("Pilih Menu : ")
		fmt.Scanln(&inputMenu)
		fmt.Println("")
		switch inputMenu {
		case 1:

			var login model.User
			fmt.Print("Username : ")
			fmt.Scanln(&login.Username)
			fmt.Print("Password :")
			fmt.Scanln(&login.Password)
			res, err := userControll.Login(login.Username, login.Password)
			if err != nil {
				fmt.Println("Erro query logn - main", err.Error())
			}
			count := len(res)
			if count == 0 {
				fmt.Println("++++ Username atau password salah ++++")
			} else {
				session = res[0]
				fmt.Println("++++ Login Berhasil! ++++")
				clear()
			}
		case 3 :
			var bukuSaya = true
			var pilih int
			if session.ID == 0 {
				bukuSaya = false
				fmt.Println("+++Silahkan login terlebih dahulu+++")
			}
			clear()
			for bukuSaya {
				// Tampilkan Semua Buku Say
				var sessId = session.ID
				res, _ := bukuControll.MyBook(sessId)
				fmt.Println("\t===== DAFTAR BUKU SAYA =====")
				fmt.Print("No")
				fmt.Print("\tKode")
				fmt.Print("\tJudul")
				fmt.Print("\tPenulis")
				fmt.Print("\tPenerbit")
				fmt.Print("\tTahun Terbit\n")

				for i := 0; i < len(res); i++ {
					fmt.Print(i+1)
					fmt.Print("\t",res[i].ID)
					fmt.Print("\t",res[i].Judul)
					fmt.Print("\t",res[i].Penulis)
					fmt.Print("\t",res[i].Penerbit)
					fmt.Print("\t",res[i].Th_terbit,"\n")
				}
				fmt.Println("")
				fmt.Println("====== Sub Menu ======")
				fmt.Println("1. Tambah Buku")
				fmt.Println("2. Edit Buku")
				fmt.Println("3. Hapus Buku")
				fmt.Println("9. Kembali")
				fmt.Print("Masukkan Pilihan : ")
				fmt.Scanln(&pilih)
				fmt.Println("")

				var buku model.Buku
				if pilih == 1 {
					fmt.Print("Judul : ")
					fmt.Scanln(&buku.Judul)
					fmt.Print("Penulis : ")
					fmt.Scanln(&buku.Penulis)
					fmt.Print("Penerbit : ")
					fmt.Scanln(&buku.Penerbit)
					fmt.Print("Tahun Terbit : ")
					fmt.Scan(&buku.Th_terbit)
					buku.ID_user = session.ID
					_, err := bukuControll.TambahBuku(buku)
					if err != nil {
						fmt.Println("Error insert buku", err.Error())
					} else {
						fmt.Println("Buku berhasil ditambahkan!")
						time.Sleep(3 * time.Second)
						clear()
					}
				} else if pilih == 2 {
					fmt.Print("Kode Buku : ")
					fmt.Scanln(&buku.ID)
					fmt.Print("Judul : ")
					fmt.Scanln(&buku.Judul)
					fmt.Print("Penulis : ")
					fmt.Scanln(&buku.Penulis)
					fmt.Print("Penerbit : ")
					fmt.Scanln(&buku.Penerbit)
					fmt.Print("Tahun Terbit : ")
					fmt.Scanln(&buku.Th_terbit)
					buku.ID_user = session.ID
					_, err := bukuControll.UpdateBuku(buku)
					if err != nil {
						fmt.Println("Error update buku", err.Error())
					} else {
						fmt.Println("Buku berhasil diupdate!")
						time.Sleep(3 * time.Second)
						clear()
					}
				} else if pilih == 3 {
						var buku model.Buku
						fmt.Print("Kode Buku : ")
						fmt.Scanln(&buku.ID)
						_, err := bukuControll.DeleteBuku(buku.ID)
						if err != nil {
							fmt.Println("Error hapus", err.Error())
						} else {
							fmt.Println("Buku berhasil dihapus!")
							time.Sleep(3 * time.Second)
							clear()
						}
				} else {
					bukuSaya = false
					clear()
				}
			}
		case 4 :
			res, err := bukuControll.GetAll()
			if err != nil {
				fmt.Print("Gagal ambil data buku", err.Error())
			}
			fmt.Print("No")
			fmt.Print("\tJudul")
			fmt.Println("\t\t\tPenulis")
			for i := 0; i < len(res); i++ {
				fmt.Print(i+1)
				fmt.Print("\t",res[i].Judul)
				fmt.Print("\t",res[i].Penulis,"\n")
			}
		}
	}
}
