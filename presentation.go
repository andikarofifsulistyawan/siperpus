package main

import "fmt"

func CetakMenuUtama(repository *BukuRepository) {
	var daftarPilihan [MEMORY_KAPASITAS_ARRAY_MAKSIMAL]string = [MEMORY_KAPASITAS_ARRAY_MAKSIMAL]string{
		"Tampilkan Semua Buku",
		"Tambah Buku",
		"Ubah Buku",
		"Hapus Buku",
		"Cari Buku",
		"Urutkan Buku",
		"Statistik Buku",
		"Keluar",
	}
	var totalPilihan int = 8
	var pilihan int

	for {
		fmt.Println("===================================================")
		fmt.Println("  SIPERPUS: SISTEM MANAJEMEN KATALOG PERPUSTAKAAN  ")
		fmt.Println("===================================================")
		pilihan = PromptPilihan("Silakan pilih salah satu menu berikut dengan memasukkan nomor pilihannya.", daftarPilihan, totalPilihan)
		fmt.Println()

		switch pilihan {
		case 1:
			CetakMenuDaftarSemuaBuku(repository)

		case 2:
			CetakMenuTambahBuku(repository)

		case 3:
			CetakMenuUbahBuku(repository)

		case 4:
			CetakMenuHapusBuku(repository)

		case 5:
			CetakMenuCariBuku(repository)

		case 6:
			CetakMenuUrutkanBuku(repository)

		case 7:
			CetakMenuStatistikBuku(repository)

		case 8:
			fmt.Println("Program selesai. Terima kasih telah menggunakan program ini!")
			return
		}

		fmt.Println()
	}
}

func CetakMenuDaftarSemuaBuku(repository *BukuRepository) {
	fmt.Println("========================")
	fmt.Println("  1. DAFTAR SEMUA BUKU  ")
	fmt.Println("========================")

	if repository.TotalBuku == 0 {
		fmt.Println("Belum ada buku.")
	} else {
		CetakKoleksiDataBuku(repository.KoleksiBuku, repository.TotalBuku)
	}

	PromptPause()
}

func CetakMenuTambahBuku(repository *BukuRepository) {
	var buku Buku
	var err error

	for {
		fmt.Println("==================")
		fmt.Println("  2. TAMBAH BUKU  ")
		fmt.Println("==================")

		buku = PromptBukuBaru()

		err = repository.CreateBuku(buku)
		if err == nil {
			fmt.Println()
			fmt.Println("Buku berhasil ditambahkan.")
		} else {
			fmt.Println()
			fmt.Println("Gagal menambahkan buku.")
			fmt.Println("Error:", err)
		}

		if !PromptBoolean("Apakah ingin menambahkan buku lainnya") {
			break
		}

		fmt.Println()
	}
}

func CetakMenuUbahBuku(repository *BukuRepository) {
	var idBukuYangDiubah string
	var indeksBukuYangDiubah int
	var bukuBaru Buku
	var err error

	for {
		fmt.Println("================")
		fmt.Println("  3. UBAH BUKU  ")
		fmt.Println("================")

		idBukuYangDiubah = PromptString("Masukkan ID buku yang ingin diubah")
		indeksBukuYangDiubah = repository.GetBukuByIDBinarySearch(idBukuYangDiubah)

		if indeksBukuYangDiubah == -1 {
			fmt.Println("Buku tidak ditemukan.")
		} else {
			fmt.Println("Buku ditemukan.")
			fmt.Println()

			CetakDataBuku(repository.KoleksiBuku[indeksBukuYangDiubah])
			fmt.Println()

			bukuBaru = PromptBukuBaru()

			err = repository.UpdateBukuByID(idBukuYangDiubah, bukuBaru)
			if err == nil {
				fmt.Println()
				fmt.Println("Buku berhasil diubah.")
			} else {
				fmt.Println()
				fmt.Println("Gagal mengubah buku.")
				fmt.Println("Error:", err)
			}
		}

		if !PromptBoolean("Apakah ingin mengubah buku lainnya") {
			break
		}

		fmt.Println()
	}
}

func CetakMenuHapusBuku(repository *BukuRepository) {
	var idBukuYangDihapus string
	var indeksBukuYangDihapus int
	var err error

	for {
		fmt.Println("=================")
		fmt.Println("  4. HAPUS BUKU  ")
		fmt.Println("=================")

		idBukuYangDihapus = PromptString("Masukkan ID buku yang ingin dihapus")
		indeksBukuYangDihapus = repository.GetBukuByIDBinarySearch(idBukuYangDihapus)

		if indeksBukuYangDihapus == -1 {
			fmt.Println("Buku tidak ditemukan.")
		} else {
			fmt.Println("Buku ditemukan.")
			fmt.Println()

			CetakDataBuku(repository.KoleksiBuku[indeksBukuYangDihapus])
			fmt.Println()

			if PromptBoolean("Apakah yakin ingin menghapusnya") {
				err = repository.DeleteBukuByID(idBukuYangDihapus)

				if err == nil {
					fmt.Println()
					fmt.Println("Buku berhasil dihapus.")
				} else {
					fmt.Println()
					fmt.Println("Gagal menghapus buku.")
					fmt.Println("Error:", err)
				}
			}
		}

		if !PromptBoolean("Apakah ingin menghapus buku lainnya") {
			break
		}

		fmt.Println()
	}
}

func CetakMenuCariBuku(repository *BukuRepository) {
	var daftarPilihan [MEMORY_KAPASITAS_ARRAY_MAKSIMAL]string = [MEMORY_KAPASITAS_ARRAY_MAKSIMAL]string{
		"Sequential Search by ID",
		"Sequential Search by Judul",
		"Binary Search by ID",
		"Binary Search by Judul",
	}
	var totalPilihan int = 4
	var pilihan int

	for {
		fmt.Println("================")
		fmt.Println("  5. CARI BUKU  ")
		fmt.Println("================")
		pilihan = PromptPilihan("Pilih salah satu menu pencarian berikut dengan memilih nomor pilihannya.", daftarPilihan, totalPilihan)

		switch pilihan {
		case 1:
			var id string = PromptString("Masukkan id buku yang ingin dicari")
			var indeksBuku int = repository.GetBukuByIDSequentialSearch(id)

			if indeksBuku != -1 {
				fmt.Println("Buku ditemukan.")
				fmt.Println()
				CetakDataBuku(repository.KoleksiBuku[indeksBuku])
			} else {
				fmt.Println("Buku tidak ditemukan.")
			}
		case 2:
			var judul string = PromptString("Masukkan judul buku yang ingin dicari")
			var indeksBuku int = repository.GetBukuByJudulSequentialSearch(judul)

			if indeksBuku != -1 {
				fmt.Println("Buku ditemukan.")
				fmt.Println()
				CetakDataBuku(repository.KoleksiBuku[indeksBuku])
			} else {
				fmt.Println("Buku tidak ditemukan.")
			}
		case 3:
			var id string = PromptString("Masukkan id buku yang ingin dicari")
			var indeksBuku int = repository.GetBukuByIDBinarySearch(id)

			if indeksBuku != -1 {
				fmt.Println("Buku ditemukan.")
				fmt.Println()
				CetakDataBuku(repository.KoleksiBuku[indeksBuku])
			} else {
				fmt.Println("Buku tidak ditemukan.")
			}
		case 4:
			var judul string = PromptString("Masukkan judul buku yang ingin dicari")
			var indeksBuku int = repository.GetBukuByJudulBinarySearch(judul)

			if indeksBuku != -1 {
				fmt.Println("Buku ditemukan.")
				fmt.Println()
				CetakDataBuku(repository.KoleksiBuku[indeksBuku])
			} else {
				fmt.Println("Buku tidak ditemukan.")
			}
		}

		if !PromptBoolean("Apakah ingin mencari lagi") {
			break
		}

		fmt.Println()
	}
}

func CetakMenuUrutkanBuku(repository *BukuRepository) {
	var daftarPilihan [MEMORY_KAPASITAS_ARRAY_MAKSIMAL]string = [MEMORY_KAPASITAS_ARRAY_MAKSIMAL]string{
		"Selection Sort by Tahun Terbit Ascending",
		"Selection Sort by Tahun Terbit Descending",
		"Insertion Sort by Tahun Terbit Ascending",
		"Insertion Sort by Tahun Terbit Descending",
	}
	var totalPilihan int = 4
	var pilihan int

	for {
		fmt.Println("===================")
		fmt.Println("  6. URUTKAN BUKU  ")
		fmt.Println("===================")
		pilihan = PromptPilihan("Pilih salah satu menu pengurutan berikut dengan memilih nomor pilihannya.", daftarPilihan, totalPilihan)

		switch pilihan {
		case 1:
			CetakKoleksiDataBuku(repository.GetAllBukuSortByTahunTerbitAscendingSelectionSort(), repository.TotalBuku)
		case 2:
			CetakKoleksiDataBuku(repository.GetAllBukuSortByTahunTerbitDescendingSelectionSort(), repository.TotalBuku)
		case 3:
			CetakKoleksiDataBuku(repository.GetAllBukuSortByTahunTerbitAscendingInsertionSort(), repository.TotalBuku)
		case 4:
			CetakKoleksiDataBuku(repository.GetAllBukuSortByTahunTerbitDescendingInsertionSort(), repository.TotalBuku)
		}

		if !PromptBoolean("Apakah ingin mengurutkan lagi") {
			break
		}

		fmt.Println()
	}
}

func CetakMenuStatistikBuku(repository *BukuRepository) {
	fmt.Println("=====================")
	fmt.Println("  7. STATISTIK BUKU  ")
	fmt.Println("=====================")

	fmt.Println("A. Total Buku per Kategori:")
	for kategori, total := range repository.TotalBukuPerKategori() {
		fmt.Printf(
			"- %s : %d buku\n",
			kategori,
			total,
		)
	}
	fmt.Println()

	fmt.Printf("B. Total Koleksi Buku: %d\n", repository.TotalBuku)
	fmt.Printf("C. Total Buku Tersedia: %d\n", repository.TotalBukuTersedia())
	PromptPause()
}

func CetakKoleksiDataBuku(koleksiBuku DaftarBuku, totalBuku int) {
	var i int

	for i = 0; i < totalBuku; i++ {
		CetakDataBuku(koleksiBuku[i])
		fmt.Println()
	}
}

func CetakDataBuku(buku Buku) {
	fmt.Printf("ID Buku: %s\n", buku.ID)
	fmt.Printf("Judul: %s\n", buku.Judul)
	fmt.Printf("Penulis: %s\n", buku.Penulis)
	fmt.Printf("Kategori: %s\n", buku.Kategori)
	fmt.Printf("Penerbit: %s\n", buku.Penerbit)
	fmt.Printf("Tahun Terbit: %d\n", buku.TahunTerbit)
	fmt.Printf("Tersedia: %t\n", buku.Tersedia)
}
