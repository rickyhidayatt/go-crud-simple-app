package console

import (
	"fmt"
	"live-code/models"
	"live-code/repository"
	"strings"
)

func PrintDosen(Id string, Nama string, Jurusan string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println("ID :", Id)
	fmt.Println("Nama :", Nama)
	fmt.Println("Jurusan :", Jurusan)
	fmt.Println()

	var exit int
	fmt.Print("Kembali ke menu utama \nKetik 0")
	fmt.Println()
	fmt.Scan(&exit)

	if exit == 0 {
		fmt.Println()
	}

}

func PrintInputDosen() {
	var dataDosens = []models.Dosen{}
	var dosenRepo = repository.DosenRepo(dataDosens)

	dataDosens = dosenRepo.DaftarDosen().DosenS
	for _, v := range dataDosens {
		PrintDosen(v.NIDN, v.Nama, v.Jurusan)
	}

	var exit int
	fmt.Print("Kembali ke menu utama \nKetik 0")
	fmt.Println()
	fmt.Scan(&exit)

	if exit == 0 {
		fmt.Println()
	}
}

func MenuDosen() {

	fmt.Println()
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println("MENU DOSEN")
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println("\n1. Tambah Dosen Baru")
	fmt.Println("2. Lihat Daftar Dosen")
	fmt.Println("3. Update Datar Dosen")
	fmt.Println("4. Hapus Daftar Dosen")
	fmt.Println("5. Exit")
	fmt.Println()

}
