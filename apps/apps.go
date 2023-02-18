package apps

import (
	"fmt"
	"live-code/entity"
)

func Run() {
	fmt.Println()
	fmt.Println("Menu : ")
	fmt.Println("1. Mahasiswa")
	fmt.Println("2. Dosen")
	fmt.Println("3. Exit")

	var input int
	fmt.Scan(&input)

	switch input {
	case 1:
		entity.MahasiswaMenuApps()
	case 2:
		fmt.Println("Menu untuk dosen")
		entity.DosenMenuApps()
	case 3:
		fmt.Println("Terimkasih telah menggunakan aplikasi kami")
		return
	default:
		fmt.Println("Kamu salah memasukan angka.")
	}
}
