package main

type Buku struct {
	ID          string
	Judul       string
	Penulis     string
	Kategori    string
	Penerbit    string
	TahunTerbit uint
	Tersedia    bool
}

type DaftarBuku [MEMORY_KAPASITAS_ARRAY_MAKSIMAL]Buku
