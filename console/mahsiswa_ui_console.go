package console

import (
	"fmt"
	"live-code/models"
	"live-code/repository"
	"strings"
)

func MenuMahasiswa() {

	fmt.Println()
	fmt.Println(strings.Repeat("=", 25))
	fmt.Println(strings.Repeat(" ", 5), "MENU MAHASISWA")
	fmt.Println(strings.Repeat("=", 25))
	fmt.Println("\n1. Tambah Mahasiswa Baru")
	fmt.Println("2. Lihat Daftar Mahasiswa")
	fmt.Println("3. Update Daftar Mahasiswa")
	fmt.Println("4. Hapus Daftar Mahasiswa")
	fmt.Println("5. Exit")
	fmt.Println("Masukan pilihan kamu :")

}

func PrintMahasiswa(Id int, Nama string, Umur int, Jurusan string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println("ID :", Id)
	fmt.Println("Nama :", Nama)
	fmt.Println("Umur :", Umur)
	fmt.Println("Jurusan :", Jurusan)
	fmt.Println()

}

func PrintInputMahasiswa() {
	var dataMahasiswaS = []models.Mahasiswa{}
	var MabaRepo = repository.MabaRepo(dataMahasiswaS)

	dataMahasiswaS = MabaRepo.DaftarMahasiswa().MahasiswaS
	for _, v := range dataMahasiswaS {
		PrintMahasiswa(v.Id, v.Nama, v.Umur, v.Jurusan)
	}

}

func Exit() {
	var exit string
	fmt.Println("Ketik Q")
	fmt.Println("Untuk Kembali ke menu utama :")
	fmt.Println()
	fmt.Scan(&exit)

	if exit == "q" || exit == "Q" {
	} else {
		fmt.Println("Kamu salah memasukan inputan")
		Exit()

	}
}
