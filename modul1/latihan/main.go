/*
penting !!!

1. string = 1 satuan, 
	var nama = "Aura" --> nama[0] isinya A

2. array = kumpulan data ukuran tetap
	var angka = [3]int{10, 20, 30} --> angka[0] isinya 10

3. slice = kumpulan data ukuran dinamis
	var angka1 = [3]int{10, 20, 30} --> angka[0] isinya 10
	var angka2 = []int{10, 20, 30} --> angka[0] isinya 10

	bedanya bisa dilakukan append(namaVar, value) 

	append di tengah :
	slice = append(slice[:index], append([]tipeData{nilaiBaru}, slice[index:]...)...)
*/


package main
 
import "fmt" // import package fmt bawaan Go untuk formatting dan I/O
 
type User struct { //class 
    ID       int    `json:"id"` // nama variabel, type data kolom, konversi ke json--> untuk bikin API
    Username string `json:"username"`
    IsActive bool   `json:"is_active"`
}
 
func (u *User) Activate()   { u.IsActive = true }
func (u User) Info() string { return fmt.Sprintf("%s aktif: %v", u.Username, u.IsActive) }

/*
u *User -> pointer receiver
           menerima pointer/alamat ke data User asli.
           Perubahan pada dilakukan di alamat data asli sehingga memengaruhi data asli.
           Cocok jika method perlu mengubah data User.
u user -> menerima salinan (copy) dari data User.
           Perubahan pada u tidak memengaruhi data asli.
           Cocok jika method hanya perlu membaca/memproses data.
*/

func tukarNilai(a, b *int) { *a, *b = *b, *a }
func main() {
    /* Variabel
	cara deklarasi variabel :
	1. var dengan tipe data dan value
		var namaVariabel tipeData = value

	2. Var tanpa value
		var namaVariabel tipe data

		Variabel akan otomatis dapat zero value sesuai tipe datanya
		contoh :
		int -> 0
		float -> 0
		string -> ""
		bool -> false

	3. var tanpa menuliskan tipe data
		var namaVariabel = value

		Go akan otomatis menentukan tipe data berdasarkan value yang diberikan
		Namun, overwrite hanya bisa dilakukan dengan value bertipe data sama. 
		Jika overwrite beda tipe data maka akan error.
	*/
    var nama string = "Sari"
    umur := 20
    var kosong string
    fmt.Printf("%s %d zero value: %q\n", nama, umur, kosong)
 
	/* Slice
	Array dengan panjang fleksibel.
	contoh :
	array = [3]int
	slice = []int 
	
	cara bikin slice :
	1. literal
		namaVar := []int(value1, value2, ...)
	2. make() --> fokus utama untuk bikin slice kosong yg baru diketahui panjang sementara
		namaVar := make([]int, panjang) --> setiap indeks diisi zero value. 
		atau
		namaVar := make([]int, )
	3, slice kosong
		namaVar := []int{}

	related function: 
	1. append() --> menambah di indeks paling belakang
		namaVar = append(var, value) 

		bisa juga klo mau tambah di tengah, 
		namaVar = append(slice[:index], append([]Typedata{value}, slice[index:]...)...)
		konsepnya masukin ke indeks sebelum add, indeks add, dan indeks setelah add
		
		untuk hapus, gaada fungsi khusus
		solusinya adalah overwrite dengan slice baru tanpa elemen yang mau dihapus
		namaVar = append(namaVar[:indekshapus], namaVar[indekshapus+1:]...)
		konsepnya masukin indeks sebelum hapus dan indeks setelah hapus. jadi indeks hapus gak dimasukin

	2. len() -> hitung panjang slice 
		len(namaVar) 

	3. cap() -> hitung kapasitas slice
		cap(namaVar)
	
	4. mengakses
		namaVar[index]

	5. mengambil sebagian (slicing)
		namaVar[indeksAwal:indeksAkhir]

	len() vs cap()
	contoh :
		s := make([]int, 3, 5)

		len = 3  → sekarang ada 3 elemen
		cap = 5  → slice punya kapasitas sampai 5 elemen

		[ 0 ][ 0 ][ 0 ][   ][   ]
  		←──── len 3 ────→
  		←────── cap 5 ─────────→
		
	
	*/
    nilai := []int{10, 20, 30}
    nilai = append(nilai, 40)
    fmt.Println(nilai, nilai[1:3], len(nilai), cap(nilai))
 
    /*MAP
	array dengan key-value. 
	syarat = key unik, value bisa sama

	deklarasi :
	namaVar := map[keyType]valueType

	process related :
	1. menambah
		namaVar[Key] = Value
	
	2. akses
		namaVar[Key]

	3. Delete 
		delete(namaVar, Key)

	4. cek ada atau tidak
		namaVarValue, namaVarBoolean := namaMap[key]
	*/

    skor := make(map[string]int)
    skor["Sari"] = 90
    if n, ada := skor["Budi"]; ada {
        fmt.Println("Budi:", n)
    } else {
        fmt.Println("Budi belum punya nilai")
    }
 
    /*CHANNEL
	media komunikasi antar goroutine

	goroutine = fungsi yang dijalankan secara concurrent oleh Go menggunakan keyword
	intinya biar bisa jalan bareng, misal fungsi A run, fungsi B run barengan 
	tidak nunggu A selesai baru B
	*/
    ch := make(chan string)
    go func() { ch <- "halo dari goroutine" }()
    fmt.Println(<-ch)
 
    /*POINTER
	untuk menunjuk alamat asli data --> dikenai proses data asli bukan salinan

	& = alamat atau indeks
	* = nilai yg berada di alamat

	Contoh :
	&value = alamat asli dari value
	*a = nilai asli variabel a 
	*/
    a, b := 1, 2
    tukarNilai(&a, &b)
    fmt.Println("setelah ditukar:", a, b)
 
    /*Struct
	mirip class di php
	*/
    u := User{ID: 1, Username: "sari"}
    u.Activate()
    fmt.Println(u.Info())
}
