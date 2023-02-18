package repository

import (
	"errors"
	"fmt"
	"live-code/models"
)

type DosenRepository interface {
	TambahDosen(maba models.Dosen)
	UpdateDosen(maba models.Dosen)
	DeleteDosen(id string) error
	DaftarDosen() dosenService
}

type dosenService struct {
	DosenS []models.Dosen
}

func (d *dosenService) TambahDosen(dosen models.Dosen) {
	d.DosenS = append(d.DosenS, dosen)
}

func (d *dosenService) DaftarDosen() dosenService {
	return *d
}

func (d *dosenService) UpdateDosen(dsns models.Dosen) {
	for i, dsn := range d.DosenS {
		if dsn.NIDN == dsn.NIDN {
			d.DosenS[i] = dsns
			break
		}
	}
}

func (dsn *dosenService) DeleteDosen(id string) error {
	// Mencari indeks dari mahasiswa dengan ID yang diberikan
	index := -1
	for i, dosens := range dsn.DosenS {
		if dosens.NIDN == id {
			index = i
			break
		}
	}

	// Jika ID tidak ditemukan, kembalikan error
	if index == -1 {
		return errors.New("Mahasiswa dengan ID tersebut tidak ditemukan")
	}

	// Menghapus mahasiswa dari slice
	dsn.DosenS = append(dsn.DosenS[:index], dsn.DosenS[index+1:]...)

	fmt.Println("Mahasiswa dengan id", id, "Berhasil dihapus")
	fmt.Println("")
	fmt.Println("Data Mahasiswa Lainya :")

	for _, v := range dsn.DosenS {
		fmt.Println()
		fmt.Println("NIDN : ", v.NIDN)
		fmt.Println("Nama : ", v.Nama)
		fmt.Println("Jurusan : ", v.Jurusan)

	}
	return nil
}

// buat manggil interface fungsi ke publik
func DosenRepo(dosen []models.Dosen) DosenRepository {
	return &dosenService{
		DosenS: dosen,
	}
}
