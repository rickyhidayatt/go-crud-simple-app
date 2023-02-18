package console

import (
	"fmt"
	"live-code/models"
	"live-code/repository"
	"strings"
)

func MenuMahasiswa() {

	fmt.Println()
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println("MENU MAHASISWA")
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println("\n1. Tambah Mahasiswa Baru")
	fmt.Println("2. Lihat Daftar Mahasiswa")
	fmt.Println("3. Update Daftar Mahasiswa")
	fmt.Println("4. Hapus Daftar Mahasiswa")
	fmt.Println("5. Exit")
	fmt.Println()

}

func PrintMahasiswa(Id int, Nama string, Umur int, Jurusan string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println("ID :", Id)
	fmt.Println("Nama :", Nama)
	fmt.Println("Umur :", Umur)
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

func PrintInputMahasiswa() {
	var dataMahasiswaS = []models.Mahasiswa{}
	var MabaRepo = repository.MabaRepo(dataMahasiswaS)

	dataMahasiswaS = MabaRepo.DaftarMahasiswa().MahasiswaS
	for _, v := range dataMahasiswaS {
		PrintMahasiswa(v.Id, v.Nama, v.Umur, v.Jurusan)
	}

	var exit int
	fmt.Print("Kembali ke menu utama \nKetik 0")
	fmt.Println()
	fmt.Scan(&exit)

	if exit == 0 {
		fmt.Println()
	}
}
