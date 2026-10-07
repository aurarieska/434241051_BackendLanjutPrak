package main

import "fmt"

// deklarasi object Student
var siswa1 = Student{
	ID : 1,
	Name : "Aura",
	Grade : 100,
	IsActive : false,
}

// deklarasi object (struct dengan value) lebih dari 1
var siswa2 = Student{
	ID:       2,
	Name:     "Sari",
	Grade:    90,
	IsActive: true,
}

var siswa3 = Student{
	ID:       3,
	Name:     "Yaya",
	Grade:    85,
	IsActive: false,
}

func main() {

	// ============================
	// NO. 2 Variabel
	// ============================
	println("string dengan tipe data dan value ", akuStringTp)
	println("string tanpa tipe data dengan value ", akuString)
	println("string tanpa tipe data dan value ", stringKosong)
	println("int dengan tipe data dan value ", iniIntTp)
	println("int tanpa tipe data dengan value ", iniInt)
	println("int tanpa tipe data dan value ", intKosyong)
	println("float dengan tipe data dan value ", itufloatTP)
	println("float tanpa tipe data dengan value ", yofloat)
	println("float tanpa tipe data dan value ", floatKosong)
	println("boolean dengan tipe data dan value ", buyiyanTp)
	println("bolean tanpa tipe data dengan value ", buliyan)
	println("boolean tanpa tipe data dan value ", boolKosong)
	fmt.Println("slice dengan tipe data dan value ", slcliteral)
	fmt.Println("slice tanpa tipe data dengan value ", slicemake)
	fmt.Println("slice tanpa tipe data dan value ", slicekosong)

	// ============================
	// NO.3 POINTER
	// ============================
	a := 10
	b := 100

	println("sebelum swapbypointer", a, b)
	swapbypointer(&a,&b)
	println(a, b)

	//buat variabel baru karena yg atas sudah kena overwrite
	c := 10
	d := 100
	println("sebelum swapbyvalue", c, d)
	swapbyvalue(c,d)
	println(c,d)

	e := []string{"aura", "rieska"}
	fmt.Println("sebelum updateslice", e)
	updateslice(&e, "maharani")
	fmt.Println("sesudah updateslice", e)


	// ============================
	// NO. 4 Struct Student
	// ============================

	/*
	Object Struct Tunggal
	*/

	//getinfo tunggal tanpa parameter karena 1 object struct
	println(GetInfo())

	// Update melalui input di luar function karena requirement soal 
	println("masukkan nilai baru")
	var nilai float64
	fmt.Scan(&nilai)
	UpdateGrade(nilai)
	println("status setelah melalui update nilai")
	println(GetInfo())

	println("status setelah melalui activate")
	activate()
	println(GetInfo())

	println("status setelah melalui deactive")
	deactive()
	println(GetInfo())

	
	/*
	Object Struct Banyak tanpa looping
	*/
	fmt.Println("=== TANPA LOOPING ===")

	// getinfo dengan parameter karena struct banyak jadi harus disebutin namaObjectnya
	fmt.Println("Informasi Siswa 1:")
	fmt.Println(GetInfoStudent(siswa1))

	fmt.Println("Informasi Siswa 2:")
	fmt.Println(GetInfoStudent(siswa2))

	fmt.Println("Informasi Siswa 3:")
	fmt.Println(GetInfoStudent(siswa3))

	fmt.Println("=== UPDATE SISWA 2 ===")

	fmt.Println("Sebelum update:")
	fmt.Println(GetInfoStudent(siswa2))

	// update melalui fixed declair dengan parameter alamat berupa &namaVariabel dan valueBaru
	UpdateGradeStudent(&siswa2, 95)

	fmt.Println("Sesudah update:")
	fmt.Println(GetInfoStudent(siswa2))

	/*
	Object Struct Banyak dengan looping
	*/

	fmt.Println("=== DENGAN LOOPING ===")

	GetAllStudent()

	//looping grade setiap object diperbarui dengan nilai berbeda
	fmt.Println("=== UPDATE NILAI SEMUA SISWA ===")

	for i := range daftarSiswa {
		var nilaiBaru float64
		
		//scan dimasukkan di dalam looping dan perbarui nilai
		fmt.Print("Masukkan nilai untuk ", daftarSiswa[i].Name, ": ")
		fmt.Scan(&nilaiBaru)
	
		UpdateGradeStudent(&daftarSiswa[i], nilaiBaru)
	}
	
	GetAllStudent()

	//looping grade setiap object diperbarui dengan nilai sama
	fmt.Println("=== UPDATE NILAI DISAMAKAN SEMUA SISWA ===")
	var nilaiBaruSama float64
	
	//nilai dimasukkan sebelum looping dan disimpan dalam variabel nilaiBaruSama
	fmt.Print("Masukkan nilai baru ")
	fmt.Scan(&nilaiBaruSama)

	//looping hanya untuk perbarui nilai setiap object dengan nilai yang sama dari nilaiBaruSama
	for i := range daftarSiswa {	
		UpdateGradeStudent(&daftarSiswa[i], nilaiBaruSama)
	}

	GetAllStudent()
}
