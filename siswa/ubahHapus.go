package siswa

import (
	"absensi/database"
	"absensi/fitur"
	"fmt"
)

func UbahHapusSiswa() {
	var nim string
	var opsi int

	fmt.Println("1. Ubah Nama Siswa")
	fmt.Println("2. Hapus Siswa")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&opsi)
	fmt.Print("Masukkan NIM target: ")
	fmt.Scan(&nim)

	var idx int = fitur.CariIndeksSeq(nim)

	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
	} else {
		if opsi == 1 {
			var namaBaru string
			fmt.Printf("Nama lama: %s. Nama baru: ", database.DataSiswa[idx].Nama)
			fmt.Scan(&namaBaru)
			database.DataSiswa[idx].Nama = namaBaru
			fmt.Println("Sukses: Nama berhasil diubah.")
		} else if opsi == 2 {
			for i := idx; i < database.NSiswa-1; i++ {
				database.DataSiswa[i] = database.DataSiswa[i+1]
			}
			database.NSiswa--
			fmt.Println("Sukses: Data siswa dihapus.")
		}
	}
}