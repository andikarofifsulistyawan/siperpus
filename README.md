# SiPerpus

SiPerpus adalah aplikasi command line interface sederhana berbasis bahasa pemrograman Go untuk mengelola koleksi buku secara sistematis. Aplikasi ini dibuat sebagai tugas besar mata kuliah Algoritma dan Pemrograman 2.

## Anggota Kelompok

| Nama | NIM | Persentase Kontribusi | Fokus Kontribusi |
|---|---|---|---|
| Andika Rofif Sulistyawan | 109082500013 | 33% | Membuat entity layer, membuat algoritma sorting: selection sort dan insertion sort; membuat logika untuk read semua data dan semua data berdasarkan pengurutan; dan membuat logika untuk menghitung total buku berdasarkan kategori serta total buku yang tersedia |
| Gilang Haryo Yudianto | 109082500021 | 33% | Membuat algoritma searching: sequential search dan binary search; membuat  logika untuk read suatu data buku tertentu, create, update, dan delete buku; dan melakukan pengujian aplikasi |
| Henda Somantri Praja | 109082500066 | 33% | Membuat seluruh presentation layer (view atau UI/UX aplikasi) dan dokumentasi aplikasi |

## Fitur

- Menambahkan data buku
- Mengubah data buku
- Menghapus data buku
- Menampilkan seluruh koleksi buku
- Mencari buku berdasarkan
  - ID buku dengan algoritma sequential search
  - ID buku dengan algoritma binary search
  - Judul buku dengan algoritma sequential search
  - Judul buku dengan algoritma binary search
- Mengurutkan koleksi buku berdasarkan tahun terbit dengan
  - Algoritma selection sort ascending
  - Algoritma selection sort descending
  - Algoritma insertion sort ascending
  - Algoritma insertion sort descending
- Menampilkan statistik:
  - Total buku per kategori
  - Total buku tersedia

## Teknologi

- Bahasa Pemrograman Go
- Command Line Interface (CLI)

## Konsep yang Digunakan

- Struct
- Alias tipe
- Array statis
- Modularitas
- Algoritma pencarian data tunggal dengan sequential search dan binary search
- Algoritma pengurutan data dengan selection sort dan insertion sort

## Cara Menjalankan Program

Pastikan Go sudah ter-install sebelum menjalankan program.

```bash
go run .
```

## Tujuan Aplikasi

- Melatih pengelolaan data menggunakan array statis
- Memahami modularitas program
- Memahami implementasi algoritma pencarian data dengan sequential search dan binary search
- Memahami implementasi algoritma pengurutan data dengan selection sort dan insertion sort
- Membangun aplikasi terminal sederhana menggunakan Go