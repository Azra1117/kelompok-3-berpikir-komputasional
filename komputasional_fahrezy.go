// SISTEM MANAJEMEN PERPUSTAKAAN SEDEARHANA (CLI) TELKOM UNIVERSITY JAKARTA
// ---------------------------------------------
// IDENTITAS KELOMPOK
// KELAS : S1TI-KJ-01
// ANGGOTA KELOMPOK :
// 1. Muhammad Fahrezy Azra (103062500068)
// 2. Raden Dimas Nugroho (103062500066)
// 3. Muhammad Naza Putra Maitara (103062500067)
// 4. Vicky Alexander (103062500041)
// 5. Habibur Rahman Faqih (103062500057)
// ---------------------------------------------
package main

import (
	"fmt"
)

// Buku adalah struktur data untuk merepresentasikan satu entitas buku
type Buku struct {
	Judul   string
	Penulis string
	Tahun   int
	Status  string // "Tersedia" atau "Dipinjam"
}

func inisialisasiBuku() []Buku {
	return []Buku{
		{"Laskar Pelangi", "Andrea Hirata", 2005, "Tersedia"},
		{"Bumi Manusia", "Pramoedya", 1980, "Tersedia"},
		{"The Great Gatsby", "F. Scott Fitzgerald", 1925, "Tersedia"},
		{"Filosofi Kopi", "Dewi Lestari", 2006, "Dipinjam"},
		{"Pengenalan Pemrograman Komputer", "Jubilee Enterprise", 2015, "Tersedia"},
		{"Pengantar 15 Bahasa Pemrograman Terbaik Di Masa Depan", "Muhammad Wali dkk", 2023, "Tersedia"},
		{"Logika Algoritma dan Pemrograman Dasar", "Ema Utami", 2019, "Dipinjam"},
		{"Dasar-dasar Teknik Informatika", "Novega Pratama", 2020, "Tersedia"},
		{"Algoritma dan Pemrograman", "Rinaldi Munir", 2018, "Tersedia"},
		{"Rekayasa Perangkat Lunak", "Roger S. Pressman", 2015, "Tersedia"},
	}
}

func tampilkanBuku(buku []Buku) {
	fmt.Println("\n--- Daftar Buku ---")
	for i, b := range buku {
		fmt.Printf("%d. %s\n", i+1, b.Judul)
		fmt.Printf("   Penulis : %s\n", b.Penulis)
		fmt.Printf("   Tahun   : %d\n", b.Tahun)
		fmt.Printf("   Status  : %s\n", b.Status)
		fmt.Println("--------------------------")
	}
}

func main() {
	buku := inisialisasiBuku()
	var menu int

	for {
		fmt.Println("\n===== SISTEM PERPUSTAKAAN SEDERHANA =====")
		fmt.Println("1. Tampilkan Daftar Buku")
		fmt.Println("2. Pinjam Buku")
		fmt.Println("3. Kembalikan Buku")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&menu)

		switch menu {

		case 1:
			tampilkanBuku(buku)

		case 2:
			tampilkanBuku(buku)
			var pilih int
			fmt.Print("Masukkan nomor buku yang ingin dipinjam: ")
			fmt.Scan(&pilih)

			if pilih < 1 || pilih > len(buku) {
				fmt.Println("Nomor buku tidak valid.")
				break
			}

			if buku[pilih-1].Status == "Tersedia" {
				buku[pilih-1].Status = "Dipinjam"
				fmt.Println("Buku berhasil dipinjam!")
			} else {
				fmt.Println("Buku sedang dipinjam.")
			}

		case 3:
			tampilkanBuku(buku)
			var pilih int
			fmt.Print("Masukkan nomor buku yang ingin dikembalikan: ")
			fmt.Scan(&pilih)

			if pilih < 1 || pilih > len(buku) {
				fmt.Println("Nomor buku tidak valid.")
				break
			}

			if buku[pilih-1].Status == "Dipinjam" {
				buku[pilih-1].Status = "Tersedia"
				fmt.Println("Buku berhasil dikembalikan!")
			} else {
				fmt.Println("Buku tidak sedang dipinjam.")
			}
		case 4:
			fmt.Println("Program selesai. Terima kasih.")
			return

		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}
