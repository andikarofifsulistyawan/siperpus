package main

import (
	"errors"
	"strings"
	"time"
)

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

func (repository *BukuRepository) GetBukuByJudulSequentialSearch(judul string) int {
	var judulLowercase string
	var i int

	judulLowercase = strings.ToLower(judul)

	for i = 0; i < repository.TotalBuku; i++ {
		if strings.ToLower(repository.KoleksiBuku[i].Judul) == judulLowercase {
			return i
		}
	}

	return -1
}

func (repository *BukuRepository) GetBukuByJudulBinarySearch(judul string) int {
	var hasil DaftarBuku
	var judulLowercase string
	var judulTengahLowercase string
	var kiri int
	var kanan int
	var tengah int

	if repository.TotalBuku == 0 {
		return -1
	}

	hasil = repository.GetAllBukuSortByJudulAscendingInsertionSort()
	judulLowercase = strings.ToLower(judul)

	kiri = 0
	kanan = repository.TotalBuku - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		judulTengahLowercase = strings.ToLower(hasil[tengah].Judul)

		if judulTengahLowercase == judulLowercase {
			return tengah
		} else if judulTengahLowercase < judulLowercase {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

func (repository *BukuRepository) CreateBuku(buku Buku) error {
	var err error

	if repository.TotalBuku == MEMORY_KAPASITAS_ARRAY_MAKSIMAL {
		return errors.New("total koleksi sudah mencapai batas maksimum")
	}

	if repository.GetBukuByIDBinarySearch(buku.ID) != -1 {
		return errors.New("ID buku sudah digunakan oleh buku lain")
	}

	err = repository.ValidasiBuku(buku)

	if err != nil {
		return err
	}

	repository.KoleksiBuku[repository.TotalBuku] = buku
	repository.TotalBuku++

	return nil
}

func (repository *BukuRepository) UpdateBukuByID(id string, buku Buku) error {
	var err error
	var index int

	index = repository.GetBukuByIDBinarySearch(id)

	if index == -1 {
		return errors.New("buku dengan ID " + id + " tidak ditemukan")
	}

	if repository.GetBukuByIDBinarySearch(buku.ID) != -1 && buku.ID != id {
		return errors.New("ID buku sudah digunakan oleh buku lain")
	}

	err = repository.ValidasiBuku(buku)

	if err != nil {
		return err
	}

	repository.KoleksiBuku[index] = buku

	return nil
}

func (repository *BukuRepository) DeleteBukuByID(id string) error {
	var index int
	var i int

	index = repository.GetBukuByIDBinarySearch(id)

	if index == -1 {
		return errors.New("buku dengan ID " + id + " tidak ditemukan")
	}

	for i = index; i < repository.TotalBuku-1; i++ {
		repository.KoleksiBuku[i] = repository.KoleksiBuku[i+1]
	}

	repository.KoleksiBuku[repository.TotalBuku-1] = Buku{}
	repository.TotalBuku--

	return nil
}

func (repository *BukuRepository) ValidasiBuku(buku Buku) error {
	if strings.TrimSpace(buku.ID) == "" {
		return errors.New("ID buku tidak boleh kosong")
	}

	if strings.TrimSpace(buku.Judul) == "" {
		return errors.New("judul buku tidak boleh kosong")
	}

	if strings.TrimSpace(buku.Penulis) == "" {
		return errors.New("penulis buku tidak boleh kosong")
	}

	if strings.TrimSpace(buku.Kategori) == "" {
		return errors.New("kategori buku tidak boleh kosong")
	}

	if strings.TrimSpace(buku.Penerbit) == "" {
		return errors.New("penerbit buku tidak boleh kosong")
	}

	if buku.TahunTerbit == 0 {
		return errors.New("tahun terbit buku harus berupa bilangan bulat positif dan tidak boleh kosong")
	}

	if buku.TahunTerbit > uint(time.Now().Year()) {
		return errors.New("tahun terbit buku tidak boleh melebihi tahun saat ini")
	}

	return nil
}

func (repository *BukuRepository) TotalBukuPerKategori() map[string]int {
	var statistikKategori map[string]int
	var i int

	statistikKategori = make(map[string]int)

	for i = 0; i < repository.TotalBuku; i++ {
		statistikKategori[repository.KoleksiBuku[i].Kategori]++
	}

	return statistikKategori
}

func (repository *BukuRepository) TotalBukuTersedia() int {
	var total int
	var i int

	total = 0

	for i = 0; i < repository.TotalBuku; i++ {
		if repository.KoleksiBuku[i].Tersedia {
			total++
		}
	}

	return total
}
