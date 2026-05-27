package main

import "fmt"

func PromptString(pertanyaan string) string {
	var jawaban string

	fmt.Printf("%s: ", pertanyaan)
	fmt.Scan(&jawaban)

	return jawaban
}

func PromptUInt(pertanyaan string) uint {
	var jawaban uint

	fmt.Printf("%s: ", pertanyaan)
	fmt.Scan(&jawaban)

	return jawaban
}

func PromptInt(pertanyaan string) int {
	var jawaban int

	fmt.Printf("%s: ", pertanyaan)
	fmt.Scan(&jawaban)

	return jawaban
}

func PromptBoolean(pertanyaan string) bool {
	var jawaban string

	for {
		fmt.Printf("%s (y/n): ", pertanyaan)
		fmt.Scan(&jawaban)

		switch jawaban {
		case "y":
			return true
		case "n":
			return false
		}

		fmt.Printf("Jawaban tidak valid. Jawab antara 'y' atau 'n'.\n")
		PromptPause()
		fmt.Println()
	}
}

func PromptPilihan(pertanyaan string, daftarPilihan [MEMORY_KAPASITAS_ARRAY_MAKSIMAL]string, totalPilihan int) int {
	var i int
	var pilihan int

	if totalPilihan == 0 {
		return 0
	}

	for {
		fmt.Println(pertanyaan)
		for i = 0; i < totalPilihan; i++ {
			fmt.Printf("%d. %s\n", i+1, daftarPilihan[i])
		}

		pilihan = PromptInt("Pilih")

		if pilihan >= 1 && pilihan <= totalPilihan {
			break
		}

		fmt.Println("Jawaban tidak valid. Pilih salah satu nomor pilihan yang ada.")
		PromptPause()
		fmt.Println()
	}

	return pilihan
}

func PromptBukuBaru() Buku {
	var buku Buku

	buku.ID = PromptString("Masukkan ID buku baru")
	buku.Judul = PromptString("Masukkan judul buku baru")
	buku.Penulis = PromptString("Masukkan penulis buku baru")
	buku.Kategori = PromptString("Masukkan kategori buku baru")
	buku.Penerbit = PromptString("Masukkan penerbit buku baru")
	buku.TahunTerbit = PromptUInt("Masukkan tahun terbit buku baru")
	buku.Tersedia = PromptBoolean("Masukkan status tersedia buku baru")

	return buku
}

func PromptPause() {
	fmt.Print("Tekan Enter untuk melanjutkan ...")
	fmt.Scanln()
	fmt.Scanln()
}