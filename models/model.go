package models

type Mahasiswa struct {
	Id      int
	Nama    string
	Umur    int
	Jurusan string
}

type Dosen struct {
	NIDN    string
	Nama    string
	Jurusan string
}
