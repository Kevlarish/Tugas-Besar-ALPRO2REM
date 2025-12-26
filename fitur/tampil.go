package fitur

import (
	"absensi/database"
	"fmt"
)

func TampilkanData(opsiSort int, ascending bool) {
	if database.NSiswa == 0 {
		fmt.Println("Data kosong.")
		return
	}

	if opsiSort == 1 {
		InsertionSortPersentase(ascending)
	} else {
		SelectionSortNIM(ascending)
	}

	fmt.Println("\n--- REKAP ABSENSI SISWA ---")
	fmt.Printf("%-5s %-15s %-20s %-10s %-10s\n", "No", "NIM", "Nama", "Jml Absen", "Hadir(%)")
	for i := 0; i < database.NSiswa; i++ {
		fmt.Printf("%-5d %-15s %-20s %-10d %-10.2f%%\n",
			i+1,
			database.DataSiswa[i].Nim,
			database.DataSiswa[i].Nama,
			database.DataSiswa[i].NRiwayat,
			database.DataSiswa[i].Persentase)
	}
}