package absensi

import "absensi/database"

func HitungStatistik(idx int) {
	var hadir, i int
	hadir = 0
	for i = 0; i < database.DataSiswa[idx].NRiwayat; i++ {
		if database.DataSiswa[idx].Riwayat[i].Status == "H" {
			hadir++
		}
	}
	if database.DataSiswa[idx].NRiwayat > 0 {
		database.DataSiswa[idx].Persentase = (float64(hadir) / float64(database.DataSiswa[idx].NRiwayat)) * 100
	} else {
		database.DataSiswa[idx].Persentase = 0.0
	}
}