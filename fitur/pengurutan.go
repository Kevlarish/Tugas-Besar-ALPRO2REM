package fitur

import (
	"absensi/database"
	"absensi/model"
)

func SelectionSortNIM(ascending bool) {
	var i, j, minMaxIdx int
	var temp model.Siswa

	for i = 0; i < database.NSiswa-1; i++ {
		minMaxIdx = i
		for j = i + 1; j < database.NSiswa; j++ {
			if ascending {
				if database.DataSiswa[j].Nim < database.DataSiswa[minMaxIdx].Nim {
					minMaxIdx = j
				}
			} else {
				if database.DataSiswa[j].Nim > database.DataSiswa[minMaxIdx].Nim {
					minMaxIdx = j
				}
			}
		}
		temp = database.DataSiswa[i]
		database.DataSiswa[i] = database.DataSiswa[minMaxIdx]
		database.DataSiswa[minMaxIdx] = temp
	}
}

func InsertionSortPersentase(ascending bool) {
	var i, j int
	var key model.Siswa
	var condition bool

	for i = 1; i < database.NSiswa; i++ {
		key = database.DataSiswa[i]
		j = i - 1
		var geser bool = true
		
		for j >= 0 && geser {
			if ascending {
				condition = database.DataSiswa[j].Persentase > key.Persentase
			} else {
				condition = database.DataSiswa[j].Persentase < key.Persentase
			}

			if condition {
				database.DataSiswa[j+1] = database.DataSiswa[j]
				j = j - 1
			} else {
				geser = false
			}
		}
		database.DataSiswa[j+1] = key
	}
}