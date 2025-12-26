package fitur

import "absensi/database"

func CariIndeksSeq(nim string) int {
	var idx int = -1
	var i int = 0
	var found bool = false
	for i < database.NSiswa && !found {
		if database.DataSiswa[i].Nim == nim {
			idx = i
			found = true
		}
		i++
	}
	return idx
}

func CariIndeksBinary(targetNim string) int {
	var left, right, mid int
	left = 0
	right = database.NSiswa - 1
	var found bool = false
	var idx int = -1

	for left <= right && !found {
		mid = (left + right) / 2
		if database.DataSiswa[mid].Nim == targetNim {
			idx = mid
			found = true
		} else if database.DataSiswa[mid].Nim < targetNim {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return idx
}