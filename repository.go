package main

type BukuRepository struct {
	KoleksiBuku DaftarBuku
	TotalBuku   int
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
