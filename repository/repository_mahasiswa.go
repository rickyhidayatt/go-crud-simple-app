package repository

import (
	"errors"
	"fmt"
	"live-code/models"
)

type MahasiswaRepository interface {
	TambahMahasiswa(maba models.Mahasiswa)
	UpdateMahasiswa(maba models.Mahasiswa)
	DeleteMahasiswa(id int) error
	DaftarMahasiswa() mahasiswaService
}

type mahasiswaService struct {
	MahasiswaS []models.Mahasiswa
}

func (m *mahasiswaService) TambahMahasiswa(maba models.Mahasiswa) {
	m.MahasiswaS = append(m.MahasiswaS, maba)
}

// balikan isi dari struct mahasismahasiswaService
func (m *mahasiswaService) DaftarMahasiswa() mahasiswaService {
	return *m
}

func (m *mahasiswaService) UpdateMahasiswa(maba models.Mahasiswa) {
	for i, mhs := range m.MahasiswaS {
		if mhs.Id == maba.Id {
			m.MahasiswaS[i] = maba
			break
		}
	}
}

func (m *mahasiswaService) DeleteMahasiswa(id int) error {
	// Mencari indeks dari mahasiswa dengan ID yang diberikan
	index := -1
	for i, maba := range m.MahasiswaS {
		if maba.Id == id {
			index = i
			break
		}
	}

	// Jika ID tidak ditemukan, kembalikan error
	if index == -1 {
		return errors.New("Mahasiswa dengan ID tersebut tidak ditemukan")
	}

	// Menghapus mahasiswa dari slice
	m.MahasiswaS = append(m.MahasiswaS[:index], m.MahasiswaS[index+1:]...)

	fmt.Println("Mahasiswa dengan id", id, "Berhasil dihapus")
	fmt.Println("")
	fmt.Println("Data Mahasiswa Lainya :")

	for _, v := range m.MahasiswaS {
		fmt.Println()
		fmt.Println("ID : ", v.Id)
		fmt.Println("Nama : ", v.Nama)
		fmt.Println("Umur : ", v.Umur)
		fmt.Println("Jurusan : ", v.Jurusan)

	}
	return nil
}

// buat manggil interface fungsi ke publik
func MabaRepo(maba []models.Mahasiswa) MahasiswaRepository {
	return &mahasiswaService{
		MahasiswaS: maba,
	}
}
