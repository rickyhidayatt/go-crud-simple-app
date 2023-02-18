package usecase

import (
	"fmt"

	"live-code/console"
	"live-code/models"
	"live-code/repository"
)

var dataDosens = []models.Dosen{}
var DosenRepo = repository.DosenRepo(dataDosens)

func InputDosen() {

	var nidn, name, jurusan string

	fmt.Print("Masukan NIDN : ")
	fmt.Scan(&nidn)
	fmt.Print("Masukan Nama Dosen : ")
	fmt.Scan(&name)
	fmt.Print("Masukan Nama Jurusan maksimal 10 Karakter : ")
	fmt.Scan(&jurusan)

	if len(name) < 3 || len(name) > 20 {
		fmt.Println("Minimal 3 sampai 20 karakter")
	} else if len(nidn) < 14 {
		fmt.Println("NIDN Minimal 14 Karakter")
	} else if len(jurusan) > 10 {
		fmt.Println("Nama jurusan terlalu panjang")
	}

	dosen := models.Dosen{NIDN: nidn, Nama: name, Jurusan: jurusan}
	DosenRepo.TambahDosen(dosen)
}

func LihatDataDosen() {
	dataDosens = DosenRepo.DaftarDosen().DosenS
	fmt.Println("Lihat Data Dosen")
	fmt.Println("1. Semua data Dosen")
	fmt.Println("2. Lihat data Dosen by ID")

	var pilihan int
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		for _, v := range dataDosens {
			console.PrintDosen(v.NIDN, v.Nama, v.Jurusan)
		}
	} else if pilihan == 2 {
		var inputnidn string
		fmt.Print("Masukan NIDN : ")
		fmt.Scan(&inputnidn)
		for _, v := range dataDosens {
			if v.NIDN == inputnidn {
				console.PrintDosen(v.NIDN, v.Nama, v.Jurusan)
			} else {
				fmt.Println()
				fmt.Println("NIDN yang kamu masukan tidak terdaftar")
				fmt.Println()
			}
		}
	}

}

func UpdateDosen() {

	var nidn, name, jurusan string

	dataDosens = DosenRepo.DaftarDosen().DosenS

	for _, v := range dataDosens {
		fmt.Print("Masukan ID Mahasiswa : ")
		fmt.Scan(&nidn)
		if nidn == v.NIDN {
			fmt.Print("Masukan Nama Mahasiswa : ")
			fmt.Scan(&name)
			fmt.Print("Masukan Nama Jurusan maksimal 10 Karakter : ")
			fmt.Scan(&jurusan)

			if len(name) < 3 || len(name) > 20 {
				fmt.Println("Minimal 3 sampai 20 karakter")
			} else if len(nidn) < 14 {
				fmt.Println("NIDN Minimal 14 Karakter")
			} else if len(jurusan) > 10 {
				fmt.Println("Nama jurusan terlalu panjang")
			}

			calonDosen := models.Dosen{NIDN: nidn, Nama: name, Jurusan: jurusan}
			DosenRepo.UpdateDosen(calonDosen)
		}
	}
	fmt.Println("Data Dosen masih kosong, silahkan input data mahasiswa terlebih dahulu")
}

func DeleteDosen() {
	var ndn string
	fmt.Print("Masukan NIDN yang ingin dihapus : ")
	fmt.Scan(&ndn)
	DosenRepo.DeleteDosen(ndn)
}
