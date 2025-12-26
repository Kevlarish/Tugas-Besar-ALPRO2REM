package absensi

import (
	"absensi/database"
	"absensi/fitur"
	"fmt"
)

func KelolaAbsensi() {
	var nim, tanggal, status string
	fmt.Print("Masukkan NIM Siswa: ")
	fmt.Scan(&nim)

	var idx int = fitur.CariIndeksSeq(nim)
	if idx == -1 {
		fmt.Println("Siswa tidak ditemukan.")
		return
	}

	fmt.Printf("Mengelola absen: %s\n", database.DataSiswa[idx].Nama)
	fmt.Print("Tanggal (DD-MM-YYYY): ")
	fmt.Scan(&tanggal)
	fmt.Print("Status (H/I/S/A): ")
	fmt.Scan(&status)

	var i int = 0
	var found bool = false
	var pos int = -1

	for i < database.DataSiswa[idx].NRiwayat && !found {
		if database.DataSiswa[idx].Riwayat[i].Tanggal == tanggal {
			pos = i
			found = true
		}
		i++
	}

	if found {
		database.DataSiswa[idx].Riwayat[pos].Status = status
		fmt.Println("Absensi diperbarui.")
	} else {
		if database.DataSiswa[idx].NRiwayat < database.NMAX {
			curr := database.DataSiswa[idx].NRiwayat
			database.DataSiswa[idx].Riwayat[curr].Tanggal = tanggal
			database.DataSiswa[idx].Riwayat[curr].Status = status
			database.DataSiswa[idx].NRiwayat++
			fmt.Println("Absensi baru dicatat.")
		} else {
			fmt.Println("Penuh.")
		}
	}
	HitungStatistik(idx)
}