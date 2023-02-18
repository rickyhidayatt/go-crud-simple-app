package usecase

import (
	"fmt"
	"live-code/console"
	"live-code/models"
	"live-code/repository"
)

var dataMahasiswaS = []models.Mahasiswa{}
var MabaRepo = repository.MabaRepo(dataMahasiswaS)

func InputMaba() {

	var umur int
	var name, jurusan string

	fmt.Print("Masukan Nama Mahasiswa : ")
	fmt.Scan(&name)
	fmt.Print("Masukan umur mahasiswa minimal 17  tahun : ")
	fmt.Scan(&umur)
	fmt.Print("Masukan Nama Jurusan maksimal 10 Karakter : ")
	fmt.Scan(&jurusan)

	if len(name) < 3 || len(name) > 20 {
		fmt.Println("Minimal 3 sampai 20 karakter")
	} else if umur < 17 {
		fmt.Println("Maaf umur anda belum cukup")
	} else if len(jurusan) > 10 {
		fmt.Println("Nama jurusan terlalu panjang")
	}

	idTambah := 1
	dataMahasiswaS = MabaRepo.DaftarMahasiswa().MahasiswaS

	for _, v := range dataMahasiswaS {
		idTambah = v.Id + 1
	}

	calonMahasiswa := models.Mahasiswa{Id: idTambah, Nama: name, Umur: umur, Jurusan: jurusan}
	MabaRepo.TambahMahasiswa(calonMahasiswa)
}

func LihatDataMaba() {
	dataMahasiswaS = MabaRepo.DaftarMahasiswa().MahasiswaS
	fmt.Println("Lihat Data Mahasiswa")
	fmt.Println("1. Semua data Mahasiswa")
	fmt.Println("2. Lihat data Mahasiswa by ID")

	var pilihan int
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		for _, v := range dataMahasiswaS {
			console.PrintMahasiswa(v.Id, v.Nama, v.Umur, v.Jurusan)
		}
	} else if pilihan == 2 {
		var inputid int
		fmt.Print("Masukan ID Mahasiswa : ")
		fmt.Scan(&inputid)
		for _, v := range dataMahasiswaS {
			if v.Id == inputid {
				console.PrintMahasiswa(v.Id, v.Nama, v.Umur, v.Jurusan)
			} else {
				fmt.Println()
				fmt.Println("ID yang kamu masukan tidak terdaftar")
				fmt.Println()
			}
		}
	}

}

func UpdateMahasiswa() {

	var id int
	var umur int
	var name, jurusan string

	dataMahasiswaS = MabaRepo.DaftarMahasiswa().MahasiswaS

	for _, v := range dataMahasiswaS {
		fmt.Print("Masukan ID Mahasiswa : ")
		fmt.Scan(&id)
		if id == v.Id {
			fmt.Print("Masukan Nama Mahasiswa : ")
			fmt.Scan(&name)
			fmt.Print("Masukan umur mahasiswa minimal 17  tahun : ")
			fmt.Scan(&umur)
			fmt.Print("Masukan Nama Jurusan maksimal 10 Karakter : ")
			fmt.Scan(&jurusan)

			if len(name) < 3 || len(name) > 20 {
				fmt.Println("Minimal 3 sampai 20 karakter")
			} else if umur < 17 {
				fmt.Println("Maaf umur anda belum cukup")
			} else if len(jurusan) > 10 {
				fmt.Println("Nama jurusan terlalu panjang")
			}

			calonMahasiswa := models.Mahasiswa{Id: id, Nama: name, Umur: umur, Jurusan: jurusan}
			MabaRepo.UpdateMahasiswa(calonMahasiswa)
		}
	}
	fmt.Println("Data Mahasiswa masih kosong, silahkan input data mahasiswa terlebih dahulu")
}

func DeleteMahasiswa() {
	var id int
	fmt.Print("Masukan ID yang ingin dihapus : ")
	fmt.Scan(&id)
	MabaRepo.DeleteMahasiswa(id)
}
