package model

const NMAX = 100

type Absensi struct {
	Tanggal string
	Status  string 
}

type Siswa struct {
	Nama       string
	Nim        string
	Riwayat    [NMAX]Absensi
	NRiwayat   int
	Persentase float64
}