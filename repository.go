package main

import "strings"

type BukuRepository struct {
	KoleksiBuku DaftarBuku
	TotalBuku   int
}

func (repository *BukuRepository) GetAllBukuSortByIDAscendingInsertionSort() DaftarBuku {
	var hasil DaftarBuku
	var temp Buku
	var i int
	var j int

	if repository.TotalBuku == 0 {
		return DaftarBuku{}
	}

	hasil = repository.KoleksiBuku

	for i = 1; i < repository.TotalBuku; i++ {
		temp = hasil[i]
		j = i

		for j > 0 && strings.ToLower(temp.ID) < strings.ToLower(hasil[j-1].ID) {
			hasil[j] = hasil[j-1]
			j--
		}

		hasil[j] = temp
	}

	return hasil
}

func (repository *BukuRepository) GetAllBukuSortByJudulAscendingInsertionSort() DaftarBuku {
	var hasil DaftarBuku
	var temp Buku
	var i int
	var j int

	if repository.TotalBuku == 0 {
		return DaftarBuku{}
	}

	hasil = repository.KoleksiBuku

	for i = 1; i < repository.TotalBuku; i++ {
		temp = hasil[i]
		j = i

		for j > 0 && strings.ToLower(temp.Judul) < strings.ToLower(hasil[j-1].Judul) {
			hasil[j] = hasil[j-1]
			j--
		}

		hasil[j] = temp
	}

	return hasil
}

func (repository *BukuRepository) GetAllBukuSortByTahunTerbitAscendingSelectionSort() DaftarBuku {
	var hasil DaftarBuku
	var temp Buku
	var indexMinimum int
	var i int
	var j int

	if repository.TotalBuku == 0 {
		return DaftarBuku{}
	}

	hasil = repository.KoleksiBuku

	for i = 0; i < repository.TotalBuku-1; i++ {
		indexMinimum = i

		for j = i + 1; j < repository.TotalBuku; j++ {
			if hasil[j].TahunTerbit < hasil[indexMinimum].TahunTerbit {
				indexMinimum = j
			}
		}

		temp = hasil[i]
		hasil[i] = hasil[indexMinimum]
		hasil[indexMinimum] = temp
	}

	return hasil
}

func (repository *BukuRepository) GetAllBukuSortByTahunTerbitDescendingSelectionSort() DaftarBuku {
	var hasil DaftarBuku
	var temp Buku
	var indexMaksimum int
	var i int
	var j int

	if repository.TotalBuku == 0 {
		return DaftarBuku{}
	}

	hasil = repository.KoleksiBuku

	for i = 0; i < repository.TotalBuku-1; i++ {
		indexMaksimum = i

		for j = i + 1; j < repository.TotalBuku; j++ {
			if hasil[j].TahunTerbit > hasil[indexMaksimum].TahunTerbit {
				indexMaksimum = j
			}
		}

		temp = hasil[i]
		hasil[i] = hasil[indexMaksimum]
		hasil[indexMaksimum] = temp
	}

	return hasil
}

func (repository *BukuRepository) GetAllBukuSortByTahunTerbitAscendingInsertionSort() DaftarBuku {
	var hasil DaftarBuku
	var temp Buku
	var i int
	var j int

	if repository.TotalBuku == 0 {
		return DaftarBuku{}
	}

	hasil = repository.KoleksiBuku

	for i = 1; i < repository.TotalBuku; i++ {
		temp = hasil[i]
		j = i

		for j > 0 && temp.TahunTerbit < hasil[j-1].TahunTerbit {
			hasil[j] = hasil[j-1]
			j--
		}

		hasil[j] = temp
	}

	return hasil
}

func (repository *BukuRepository) GetAllBukuSortByTahunTerbitDescendingInsertionSort() DaftarBuku {
	var hasil DaftarBuku
	var temp Buku
	var i int
	var j int

	if repository.TotalBuku == 0 {
		return DaftarBuku{}
	}

	hasil = repository.KoleksiBuku

	for i = 1; i < repository.TotalBuku; i++ {
		temp = hasil[i]
		j = i

		for j > 0 && temp.TahunTerbit > hasil[j-1].TahunTerbit {
			hasil[j] = hasil[j-1]
			j--
		}

		hasil[j] = temp
	}

	return hasil
}

func (repository *BukuRepository) GetBukuByIDSequentialSearch(id string) int {
	var idLowercase string
	var i int

	idLowercase = strings.ToLower(id)

	for i = 0; i < repository.TotalBuku; i++ {
		if strings.ToLower(repository.KoleksiBuku[i].ID) == idLowercase {
			return i
		}
	}

	return -1
}

func (repository *BukuRepository) GetBukuByIDBinarySearch(id string) int {
	var hasil DaftarBuku
	var idLowercase string
	var idTengahLowercase string
	var kiri int
	var kanan int
	var tengah int

	if repository.TotalBuku == 0 {
		return -1
	}

	hasil = repository.GetAllBukuSortByIDAscendingInsertionSort()
	idLowercase = strings.ToLower(id)

	kiri = 0
	kanan = repository.TotalBuku - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		idTengahLowercase = strings.ToLower(hasil[tengah].ID)

		if idTengahLowercase == idLowercase {
			return tengah
		} else if idTengahLowercase < idLowercase {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}
