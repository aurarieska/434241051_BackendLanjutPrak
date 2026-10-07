package main

import "fmt"

func cobaubah() {
	// Go otomatis menentukan tipe data berdasarkan value
	var umur = 20

	fmt.Println("Nilai awal:", umur)

	// Value boleh diubah karena 25 masih bertipe int
	umur = 25

	// over write eror dengan timpa beda type data
	// umur = "dua puyuhyima"

	fmt.Println("Setelah diubah:", umur)

	// Baris berikut TIDAK BISA dijalankan karena
	// umur sudah ditentukan sebagai int.
	// umur = "dua puluh"

	fmt.Printf("Tipe data umur: %T\n", umur)
}