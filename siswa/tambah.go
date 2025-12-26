package siswa

import (
	"absensi/database"
	"absensi/fitur" 
	"fmt"
)

func TambahSiswa() {
	if database.NSiswa >= database.NMAX {
		fmt.Println("Gagal: Kapasitas array penuh.")
		return
	}

	var nim, nama string
	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nim)

	
	if fitur.CariIndeksSeq(nim) != -1 {
		fmt.Println("Gagal: NIM sudah terdaftar.")
	} else {
		fmt.Print("Masukkan Nama (Gunakan_Underscore): ")
		fmt.Scan(&nama)


		database.DataSiswa[database.NSiswa].Nim = nim
		database.DataSiswa[database.NSiswa].Nama = nama
		database.DataSiswa[database.NSiswa].NRiwayat = 0
		database.DataSiswa[database.NSiswa].Persentase = 0.0
		database.NSiswa++
		fmt.Println("Sukses: Data siswa berhasil ditambahkan.")
	}
}