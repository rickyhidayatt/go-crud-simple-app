package entity

import (
	"fmt"
	"live-code/console"
	"live-code/usecase"
)

func DosenMenuApps() {

	console.MenuDosen()

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
				usecase.InputDosen()
				console.PrintInputDosen()
				console.MenuDosen()
			case 2:
				usecase.LihatDataDosen()
				console.MenuDosen()
			case 3:
				usecase.UpdateDosen()
				console.MenuDosen()
			case 4:
				usecase.DeleteDosen()
				console.MenuDosen()

			default:
				fmt.Println("Pilihan yang anda masukan salah")
				console.MenuDosen()
			}
		}
	}
}
