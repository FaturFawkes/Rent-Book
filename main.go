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
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Buku{})
	db.AutoMigrate(&model.Rent{})

}

func conn() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/rent-book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// func Logout() {
// 	cmd := exec.Command("clear")
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()
// }

// func Relog() {
// 	cmd := exec.Command("go run .")
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()
// }

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

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	Clear()
	var isRunning bool = true
	conn, err := conn()
	if err != nil {
		fmt.Println("error", err.Error())
	}

	// migrate(con

	// SET MODEL CONTROLLER
	session := model.User{}
	userModel := model.UserModel{conn}
	userControll := controller.UserController{userModel}
	bukuModel := model.BukuModel{conn}
	bukuControll := controller.BukuController{bukuModel}
	rentModel := model.RentModel{conn}
	rentControll := controller.RentController{rentModel}

	for isRunning {
		var inputMenu int
		fmt.Println("")
		fmt.Println("===== RENT BOOK APP =====")
		fmt.Println("")
		fmt.Println("1.Login")
		fmt.Println("2.Update Profile(Login)")
		fmt.Println("3.Hapus Profile(Login)")
		fmt.Println("4.Buku Saya (Login)")
		fmt.Println("5.Daftar Buku")
		fmt.Println("6.Pinjam Buku (Login)")
		fmt.Println("7.Register")
		fmt.Println("8.Logout")
		fmt.Println("")
		fmt.Print("Pilih Menu : ")
		fmt.Scanln(&inputMenu)
		fmt.Println("")
		switch inputMenu {

		//LOGIN
		case 1:
			var login model.User
			fmt.Print("Username : ")
			fmt.Scanln(&login.Username)
			fmt.Print("Password : ")
			fmt.Scanln(&login.Password)
			res, _ := userControll.Login(login.Username, login.Password)
			count := len(res)
			if count == 0 {
				fmt.Println("++++ Username atau password salah ++++")
				time.Sleep(2 * time.Second)
				Clear()
			} else {
				session = res[0]
				fmt.Println("++++ Login Berhasil! ++++")
				time.Sleep(2 * time.Second)
				Clear()
			}

		//UPDATE USER
		case 2:
			var update model.User
			var cekLogin = true
			if session.ID == 0 {
				cekLogin = false
				fmt.Println("++++ Silahkan login terlebih dahulu ++++")
				time.Sleep(3 * time.Second)
				Clear()
			}
			if cekLogin == true {
				fmt.Print("Nama : ")
				fmt.Scanln(&update.Nama)
				fmt.Print("Username : ")
				fmt.Scanln(&update.Username)
				fmt.Print("Password : ")
				fmt.Scanln(&update.Password)
				fmt.Print("Email : ")
				fmt.Scanln(&update.Email)
				fmt.Print("Alamat : ")
				fmt.Scanln(&update.Alamat)
				update.ID = session.ID
				_, err := userControll.Update(update)
				if err != nil {
					fmt.Println("gagal update")
				} else {
					fmt.Println("Update berhasil!")
					time.Sleep(3 * time.Second)
					Clear()
				}
			}
			// DELETE USER
		case 3:
			var Delete model.User
			var cekDelete = true
			if session.ID == 0 {
				cekDelete = false
				fmt.Println("++++ Silahkan login terlebih dahulu ++++")
				time.Sleep(3 * time.Second)
				Clear()
			}
			if cekDelete == true {
				Delete.ID = session.ID
				_, err := userControll.Delete(session)
				if err != nil {
					fmt.Println("gagal Delete")
				} else {
					fmt.Println("Delete Berhasil")
					session.ID = 0
					time.Sleep(3 * time.Second)
					Clear()
				}
			}

		// BUKU SAYA
		case 4:
			var bukuSaya = true
			var pilih int
			if session.ID == 0 {
				bukuSaya = false
				fmt.Println("++++ Silahkan login terlebih dahulu ++++")
				time.Sleep(3 * time.Second)
				Clear()
			}
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
					fmt.Print(i + 1)
					fmt.Print("\t", res[i].ID)
					fmt.Print("\t", res[i].Judul)
					fmt.Print("\t", res[i].Penulis)
					fmt.Print("\t", res[i].Penerbit)
					fmt.Print("\t", res[i].Th_terbit, "\n")
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
						pilih = 9
						time.Sleep(3 * time.Second)
						Clear()
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
						pilih = 9
						time.Sleep(3 * time.Second)
						Clear()
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
						pilih = 9
						time.Sleep(3 * time.Second)
						Clear()
					}
				} else if pilih == 9 {
					bukuSaya = false
					Clear()
				}
			}

		// DAFTAR BUKU
		case 5:
			res, err := bukuControll.GetAll()
			if err != nil {
				fmt.Print("Gagal ambil data buku", err.Error())
			}
			fmt.Println("\t===== DAFTAR SELURUH BUKU =====")
			fmt.Println()
			fmt.Print("No")
			fmt.Print("\tKode")
			fmt.Print("\tPenulis")
			fmt.Print("\tPenerbit")
			fmt.Print("\tJudul")
			fmt.Println("\tPemilik")
			for i := 0; i < len(res); i++ {
				fmt.Print(i + 1)
				fmt.Print("\t", res[i].Id)
				fmt.Print("\t", res[i].Penulis)
				fmt.Print("\t", res[i].Penerbit)
				fmt.Print("\t", res[i].Judul)  //
				fmt.Println("\t", res[i].Nama) //
			}

			// PINJAM BUKU
		case 6:
			var bukuRent = true
			if session.ID == 0 {
				bukuRent = false
				fmt.Println("++++ Silahkan login terlebih dahulu ++++")
				time.Sleep(3 * time.Second)
				Clear()
			}
			for bukuRent {
				var pilih int
				res, _ := rentControll.GetAll()
				fmt.Println("\t===== DAFTAR BUKU SEDANG DIPINJAM =====")
				fmt.Println()
				fmt.Print("No")
				fmt.Print("\tKode")
				fmt.Print("\tJudul")
				fmt.Print("\t\tPemilik\n")
				for i := 0; i < len(res); i++ {
					fmt.Print(i + 1)
          			fmt.Print("\t", res[i].Id)
					fmt.Print("\t", res[i].Judul)
					fmt.Println("\t", res[i].Nama)
				}
				fmt.Println("")
				fmt.Println("====== Sub Menu ======")
				fmt.Println("1. Pinjam Buku")
				fmt.Println("2. Pengembalian Buku")
				fmt.Println("9. Kembali")
				fmt.Print("Masukkan Pilihan : ")
				fmt.Scanln(&pilih)
				fmt.Println("")

				if pilih == 1 {
					var book uint
					fmt.Print("Kode Buku : ")
					fmt.Scanln(&book)
					fmt.Println("")
					// CEK BUKU SUDAH DIPINJAM ?
					res := rentControll.CekRent(book)
					if res != true {
						fmt.Println("Buku sudah dipinjam!")
						fmt.Println("")
						time.Sleep(2 * time.Second)
						Clear()
					} else {
						userId := session.ID
						_, err := rentControll.AddRent(book, userId)
						if err != nil {
							fmt.Println("Gagal input rent")
						}
						fmt.Println("Buku berhasil dipinjam!")
						time.Sleep(2 * time.Second)
						Clear()
					}
				} else if pilih == 2 {
					// PENGEMBALIAN BUKU
					var kode uint
					fmt.Println("")
					fmt.Print("Kode Buku : ")
					fmt.Scanln(&kode)
					fmt.Println("")
					res, err := rentControll.KembaliBuku(kode, session.ID)
					if err != nil {
						fmt.Println("Pengembalian gagal!")
					}
					if res == true {
						fmt.Println("Pengembalian Berhasil!")
						time.Sleep(2 * time.Second)
						Clear()
					}
				} else {
					bukuRent = false
						Clear()
				}
				}

			// UPDATE USER
		case 7:
			var user model.User
			fmt.Print("Nama : ")
			fmt.Scanln(&user.Nama)
			fmt.Print("Username : ")
			fmt.Scanln(&user.Username)
			fmt.Print("Password : ")
			fmt.Scanln(&user.Password)
			fmt.Print("Email : ")
			fmt.Scanln(&user.Email)
			fmt.Print("Alamat : ")
			fmt.Scanln(&user.Alamat)
			_, err := userControll.Insert(user)
			if err != nil {
				fmt.Println("gagal register")
			} else {

				fmt.Println("Registrasi user berhasil")
				time.Sleep(3 * time.Second)
				Clear()
			}

		// LOGOUT
		case 8:
			session = model.User{}
			fmt.Println("Terima Kasih")
			time.Sleep(2 * time.Second)
			Clear()
		default:
			Clear()
		}
	}
}
