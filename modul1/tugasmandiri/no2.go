package main // 2. Deklarasi lima variabel dengan tipe berbeda: string, int, float64, bool, dan slice

// STRING
var akuStringTp string = "Aura Rieska Maharani"
var akuString = "Aura Rieska Maharani" 
// hasil aja declrasi maupun tidak deklarasi tipe data, bedanya klo type data dideklarasi sejak awal maka akan strict dan menolak value yg bukan string
// eks : var akuString String = 20 --> ditolak karena integer, tidak sesuai tipe data yang dideklarasikan. 
var stringKosong string // zero value string = "", harus sebutin tipe data agar tau zero valuenya

// INT
var iniIntTp int = 20
var iniInt = 21
var intKosyong int

// FLOAT
var itufloatTP float64 = 23.923892839
var yofloat = 28.30290
var floatKosong float64

// BOOLEAN
var buyiyanTp bool = true
var buliyan = false
var boolKosong bool

// SLICE
// slice literal
var slcliteral = []int{1, 33, 2323, 433, 32, 445}

// make zero value --> menyebutkan panjang awal, semua diisi zero value
var slicemake = make([]int, 7)

// slice kosong --> di awal 1 ruang diisi zero value
var slicekosong = []int{}















