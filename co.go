package main

import (
	"fmt"
)
type Resep struct {
	Judul          string
	BahanUtama     string // Bahan pencarian utama
	DaftarBahan    [20]string
	JumlahBahan    int
	LangkahLangkah [20]string
	JumlahLangkah  int
	DurasiMasak    int
	Kategori       string
	JumlahDicari   int
}
var daftarResep [100]Resep
var jumlahData int = 0
func main() {
	var pilihan int
	for pilihan != 9 { 
		fmt.Println("\n=== ResepKu: Manajemen Kuliner ===")
		fmt.Println("1. Tambah Resep Baru")
		fmt.Println("2. Urutkan Berdasarkan Durasi (Selection Sort)")
		fmt.Println("3. Urutkan Berdasarkan Judul (Insertion Sort)")
		fmt.Println("4. Cari Berdasarkan Bahan Utama (Sequential Search)")
		fmt.Println("5. Cari Berdasarkan Judul (Binary Search - Harus diurutkan dulu!)")
		fmt.Println("6. Statistik Kategori")
		fmt.Println("7. Edit Resep")
		fmt.Println("8. Hapus Resep")
		fmt.Println("9. Keluar")
// ... sesuaikan logic if/else untuk menu 7, 8, dan 9
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		
		if pilihan == 1 {
			tambahResepInteraktif()
		} else if pilihan == 2 {
			selectionSortDurasi()
			tampilkanDataSingkat()
		} else if pilihan == 3 {
			insertionSortJudul()
			tampilkanDataSingkat()
		} else if pilihan == 4 {
			var cari string
			fmt.Print("Masukkan Bahan Utama yang dicari: ")
			fmt.Scan(&cari)
			sequentialSearch(cari)
		} else if pilihan == 5 {
			var cariJudul string
			fmt.Print("Masukkan Judul Resep yang dicari: ")
			fmt.Scan(&cariJudul)
			binarySearchJudul(cariJudul)
		} else if pilihan == 6 {
			tampilkanStatistik()
		}else if pilihan == 7 {
			editResep()
		}else if pilihan == 8 {
			hapusResep()
		}
	}
}
func tambahResepInteraktif() {
	if jumlahData >= 100 {
		fmt.Println("Penyimpanan penuh!")
		return
	}
	var r Resep
	var input string

	fmt.Print("Judul Resep: ")
	fmt.Scan(&r.Judul)
	fmt.Print("Bahan Utama (untuk pencarian): ")
	fmt.Scan(&r.BahanUtama)
	fmt.Println("Masukkan Detail Bahan (ketik 'selesai' untuk berhenti):")
	for i := 0; i < 20; i++ {
		fmt.Printf("Bahan %d: ", i+1)
		fmt.Scan(&input)
		if input == "selesai" {
			break
		}
		r.DaftarBahan[i] = input
		r.JumlahBahan++
	}
	fmt.Println("Masukkan Langkah Memasak (ketik 'selesai' untuk berhenti):")
	for i := 0; i < 20; i++ {
		fmt.Printf("Langkah %d: ", i+1)
		fmt.Scan(&input)
		if input == "selesai" {
			break
		}
		r.LangkahLangkah[i] = input
		r.JumlahLangkah++
	}
	fmt.Println("Estimasi Durasi (menit): ")
	fmt.Scan(&r.DurasiMasak)
	fmt.Print("Kategori (Daging/Sayur/dll): ")
	fmt.Scan(&r.Kategori)
	daftarResep[jumlahData] = r
	jumlahData++
	fmt.Println("Resep berhasil ditambahkan!")
}
func selectionSortDurasi() {
	for i := 0; i < jumlahData-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahData; j++ {
			if daftarResep[j].DurasiMasak < daftarResep[minIdx].DurasiMasak {
				minIdx = j
			}
		}
		temp := daftarResep[i]
		daftarResep[i] = daftarResep[minIdx]
		daftarResep[minIdx] = temp
	}
	fmt.Println("\nData berhasil diurutkan berdasarkan Durasi!")
}

func editResep() {
	tampilkanDataSingkat()
	var idx int
	fmt.Print("Masukkan nomor resep yang ingin diedit: ")
	fmt.Scan(&idx)
	idx-- // Sesuaikan dengan indeks array (0-based)

	if idx >= 0 && idx < jumlahData {
		fmt.Printf("Mengedit: %s\n", daftarResep[idx].Judul)
		// Ulangi proses input untuk data yang ingin diubah
		fmt.Print("Judul baru: ")
		fmt.Scan(&daftarResep[idx].Judul)
		fmt.Print("Durasi baru (menit): ")
		fmt.Scan(&daftarResep[idx].DurasiMasak)
		fmt.Println("Resep berhasil diperbarui!")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func hapusResep() {
	tampilkanDataSingkat()
	var idx int
	fmt.Print("Masukkan nomor resep yang ingin dihapus: ")
	fmt.Scan(&idx)
	idx--

	if idx >= 0 && idx < jumlahData {
		// Geser elemen setelah yang dihapus ke kiri
		for i := idx; i < jumlahData-1; i++ {
			daftarResep[i] = daftarResep[i+1]
		}
		jumlahData--
		fmt.Println("Resep berhasil dihapus!")
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func insertionSortJudul() {
	for i := 1; i < jumlahData; i++ {
		key := daftarResep[i]
		j := i - 1
		for j >= 0 && daftarResep[j].Judul > key.Judul {
			daftarResep[j+1] = daftarResep[j]
			j = j - 1
		}
		daftarResep[j+1] = key
	}
	fmt.Println("\nData berhasil diurutkan berdasarkan Judul!")
}

func sequentialSearch(target string) {
	ketemu := false
	for i := 0; i < jumlahData; i++ {
		if daftarResep[i].BahanUtama == target {
			tampilkanDetailResep(i)
			daftarResep[i].JumlahDicari++
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Resep dengan bahan tersebut tidak ditemukan.")
	}
}
func binarySearchJudul(target string) {
	low := 0
	high := jumlahData - 1
	ketemu := false

	for low <= high {
		mid := (low + high) / 2

		if daftarResep[mid].Judul == target {
			tampilkanDetailResep(mid)
			daftarResep[mid].JumlahDicari++
			ketemu = true
			break // Keluar dari loop jika sudah ketemu
		} else if daftarResep[mid].Judul < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !ketemu {
		fmt.Println("Resep dengan judul tersebut tidak ditemukan.")
		fmt.Println("(Catatan: Pastikan Anda sudah menggunakan menu 3 sebelum mencari dengan Binary Search!)")
	}
}

func tampilkanDetailResep(idx int) {
	fmt.Printf("\n--- %s ---\n", daftarResep[idx].Judul)
	fmt.Printf("Kategori: %s\n", daftarResep[idx].Kategori)
	fmt.Printf("Durasi: %d menit\n", daftarResep[idx].DurasiMasak)
	fmt.Println("Bahan-bahan:")
	for b := 0; b < daftarResep[idx].JumlahBahan; b++ {
		fmt.Printf("  - %s\n", daftarResep[idx].DaftarBahan[b])
	}
	fmt.Println("Langkah:")
	for l := 0; l < daftarResep[idx].JumlahLangkah; l++ {
		fmt.Printf("  %d. %s\n", l+1, daftarResep[idx].LangkahLangkah[l])
	}
}

func tampilkanStatistik() {
	for i := 0; i < jumlahData; i++ {
		pernah := false
		for k := 0; k < i; k++ {
			if daftarResep[i].Kategori == daftarResep[k].Kategori {
				pernah = true
			}
		}
		if !pernah {
			count := 0
			for j := 0; j < jumlahData; j++ {
				if daftarResep[j].Kategori == daftarResep[i].Kategori {
					count++
				}
			}
			fmt.Printf("Kategori %s: %d resep\n", daftarResep[i].Kategori, count)
			
		}

	}
}
func tampilkanDataSingkat() {
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("%d. %s (%d mnt) - Kategori: %s\n", i+1, daftarResep[i].Judul, daftarResep[i].DurasiMasak, daftarResep[i].Kategori)
	}
}



