package entity

import (
	"fmt"
	"live-code/console"
	"live-code/usecase"
)

func MahasiswaMenuApps() {

	console.MenuMahasiswa()

	var pilih int

	for {
		fmt.Scan(&pilih)

		if pilih == 5 {
			fmt.Println("Terimkasih sudah menggunakan aplikasi")
			break
			//stop balik ke menu awal
		} else {
			//menu app mahasiswa
			switch pilih {
			case 1:
				usecase.InputMaba()
				console.PrintInputMahasiswa()
				console.MenuMahasiswa()
			case 2:
				usecase.LihatDataMaba()
				console.MenuMahasiswa()
			case 3:
				usecase.UpdateMahasiswa()
				console.MenuMahasiswa()
			case 4:
				usecase.DeleteMahasiswa()
				console.MenuMahasiswa()

			default:
				fmt.Println("Pilihan yang anda masukan salah")
				console.MenuMahasiswa()
			}
		}
	}
}
