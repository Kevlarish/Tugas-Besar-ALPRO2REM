package main

import (
	"absensi/absensi"
	"absensi/fitur"
	"absensi/siswa"
	"fmt"
)

func main() {
	var pilihan int
	var selesai bool = false

	for !selesai {
		fmt.Println("\n=== APLIKASI ABSENSI SISWA ===")
		fmt.Println("1. Tambah Siswa")
		fmt.Println("2. Kelola Absensi")
		fmt.Println("3. Ubah/Hapus Siswa")
		fmt.Println("4. Cari Siswa (Binary Search)")
		fmt.Println("5. Tampilkan Data (Sorting)")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			siswa.TambahSiswa()
		case 2:
			absensi.KelolaAbsensi()
		case 3:
			siswa.UbahHapusSiswa()
		case 4:
			var nim string
			fmt.Print("Masukkan NIM: ")
			fmt.Scan(&nim)

			fitur.SelectionSortNIM(true) 
			idx := fitur.CariIndeksBinary(nim)
			if idx != -1 {
				fmt.Println("Data Ditemukan!")
			} else {
				fmt.Println("Tidak ditemukan.")
			}
		case 5:
			var mode, urut int
			fmt.Println("1. By Persentase\n2. By NIM")
			fmt.Print("Pilih: ")
			fmt.Scan(&mode)
			fmt.Println("1. Ascending\n2. Descending")
			fmt.Print("Pilih: ")
			fmt.Scan(&urut)
			fitur.TampilkanData(mode, urut == 1)
		case 6:
			selesai = true
		}
	}
}