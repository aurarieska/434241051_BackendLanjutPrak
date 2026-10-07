penggunaan ai digunakan untuk mempelajari skenario alternatif yang tidak tercantum pada requirement soal dalam mencapai pemahaman, serta memahami command baru.

# AI Usage

## 1. Penggunaan AI

Dalam pengerjaan modul Go, AI digunakan sebagai alat bantu pembelajaran, pemahaman konsep, debugging, dan pengecekan kode. Penggunaan AI tidak dimaksudkan untuk menggantikan proses pengerjaan, tetapi untuk membantu memahami konsep serta mencari penyebab kesalahan pada kode yang dibuat.

AI yang digunakan adalah **ChatGPT**.

---

## 2. Penggunaan AI pada Room Chat Modul Go

Pada room chat ini, AI digunakan untuk membantu memahami dan mengembangkan materi Go, khususnya pada beberapa topik berikut.

### 2.1 Package dan `package main`

AI digunakan untuk memahami fungsi `package main`, hubungan antar-file dalam package yang sama, serta cara menjalankan beberapa file Go dalam satu package.

Contoh pembahasan:

* Perbedaan `package main` dengan nama module.
* Hubungan `main.go` dengan file Go lainnya seperti `no4a.go`.
* Penggunaan beberapa file dengan `package main`.
* Perbedaan menjalankan `go run main.go` dengan `go run .`.

### 2.2 Deklarasi Variabel

AI digunakan untuk memahami beberapa bentuk deklarasi variabel pada Go, yaitu:

```go
var nama string = "Aura"
```

```go
var nama = "Aura"
```

```go
nama := "Aura"
```

Pembahasan juga mencakup:

* Perbedaan `var` dan `:=`.
* Tipe data yang ditentukan secara otomatis oleh Go.
* Zero value.
* Batasan penggunaan `:=` pada package scope dan function scope.
* Apakah tipe data variabel dapat berubah setelah dideklarasikan.

### 2.3 Package `fmt`

AI digunakan untuk memahami fungsi-fungsi dari package `fmt`, antara lain:

```go
fmt.Print()
fmt.Println()
fmt.Printf()
fmt.Sprint()
fmt.Sprintln()
fmt.Sprintf()
```

Pembahasan juga mencakup penggunaan format seperti:

```go
%d
%v
%T
%s
%f
```

### 2.4 Pointer

AI digunakan untuk memahami konsep pointer, terutama penggunaan:

```go
&
```

dan

```go
*
```

Pembahasan meliputi:

* `&` sebagai operator untuk mendapatkan alamat suatu variabel.
* `*` sebagai tipe pointer dan dereference operator.
* Perbedaan parameter value dan parameter pointer.
* Mengapa perubahan melalui pointer dapat mengubah data asli.
* Contoh penggunaan pointer pada function.

Contoh yang dipelajari:

```go
func UpdateGradeStudent(siswa *Student, grade float64) {
    siswa.Grade = grade
}
```

dan pemanggilannya:

```go
UpdateGradeStudent(&siswa2, 95)
```

### 2.5 Struct

AI digunakan untuk membantu memahami penggunaan `struct` sebagai tipe data bentukan.

Contoh:

```go
type Student struct {
    ID       int
    Name     string
    Grade    float64
    IsActive bool
}
```

Pembahasan meliputi:

* Deklarasi `struct`.
* Pembuatan object dari `struct`.
* Pengisian field.
* Pengaksesan field menggunakan `.`
* Perbedaan object `struct` tunggal dengan beberapa object.
* Penggunaan struct sebagai parameter function.
* Penggunaan pointer ke struct.
* Penggunaan struct tag untuk JSON.

Contoh object:

```go
var siswa1 = Student{
    ID:       1,
    Name:     "Aura",
    Grade:    100,
    IsActive: false,
}
```

### 2.6 Function dengan Parameter

AI digunakan untuk memahami alasan penggunaan parameter pada function.

Contoh:

```go
func GetInfoStudent(siswa Student) string {
    return fmt.Sprintln(
        siswa.ID,
        siswa.Name,
        siswa.Grade,
        siswa.IsActive,
    )
}
```

Function tersebut kemudian dipanggil dengan object tertentu:

```go
GetInfoStudent(siswa1)
GetInfoStudent(siswa2)
GetInfoStudent(siswa3)
```

Hal ini digunakan untuk memahami perbedaan antara function yang langsung menggunakan object tertentu dengan function yang dapat menerima object berbeda melalui parameter.

### 2.7 Scope Variabel

AI digunakan untuk memahami perbedaan antara variabel yang dideklarasikan pada package scope dan function scope.

Contoh package scope:

```go
var siswa1 = Student{
    ID: 1,
    Name: "Aura",
}
```

Contoh function scope:

```go
func main() {
    var siswa1 = Student{
        ID: 1,
        Name: "Aura",
    }
}
```

Pembahasan digunakan untuk memahami mengapa function yang berada pada file lain masih dapat mengakses `siswa1` apabila object tersebut berada pada package scope dan kedua file menggunakan:

```go
package main
```

### 2.8 Struct dengan Beberapa Object

AI digunakan untuk memahami pembuatan beberapa object dari satu `struct`, misalnya:

```go
var siswa1 = Student{
    ID: 1,
    Name: "Aura",
}

var siswa2 = Student{
    ID: 2,
    Name: "Sari",
}

var siswa3 = Student{
    ID: 3,
    Name: "Yaya",
}
```

Pembahasan dilakukan untuk memahami bahwa `siswa1`, `siswa2`, dan `siswa3` merupakan object berbeda dengan tipe yang sama, yaitu `Student`.

### 2.9 Slice dan Looping

AI juga digunakan untuk memahami percobaan penyimpanan beberapa object struct ke dalam slice:

```go
var daftarSiswa = []Student{
    siswa1,
    siswa2,
    siswa3,
}
```

Kemudian data ditampilkan menggunakan looping:

```go
for i, siswa := range daftarSiswa {
    fmt.Println("Student ke-", i+1)
    fmt.Println("ID:", siswa.ID)
    fmt.Println("Nama:", siswa.Name)
    fmt.Println("Grade:", siswa.Grade)
    fmt.Println("Status:", siswa.IsActive)
}
```

Pembahasan digunakan untuk memahami hubungan antara `struct`, object, slice, dan looping.

---

## 3. Penggunaan AI pada Room Chat Lain

Selain room chat ini, AI juga digunakan pada room chat lain yang masih berkaitan dengan pembelajaran dan pengerjaan modul Go.

Beberapa pembahasan yang dilakukan antara lain:

### 3.1 Setup Go dan Fiber

AI digunakan untuk membantu proses awal penggunaan Go, termasuk:

* Pemeriksaan versi Go.
* Pembuatan module menggunakan `go mod init`.
* Instalasi dependency.
* Penggunaan Fiber.
* Struktur project Go.
* Pembuatan server sederhana.
* Penggunaan `app.Listen()`.

Contoh kode yang dipelajari:

```go
app := fiber.New()

app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
})

log.Fatal(app.Listen(":3000"))
```

### 3.2 HTTP Status dan Testing API

AI digunakan untuk membantu memahami response HTTP pada aplikasi Fiber, termasuk status seperti:

```text
201 Created
200 OK
422 Unprocessable Entity
415 Unsupported Media Type
204 No Content
```

AI juga digunakan untuk membantu memahami pengujian endpoint menggunakan command line.

### 3.3 Git

AI digunakan sebagai bantuan dalam penggunaan Git untuk project Go, antara lain:

```bash
git init
git add .
git commit
git push
```

serta memahami operasi seperti:

* Mengubah nama.
* Membuat commit.
* Menghapus commit terakhir.
* Memperbarui repository setelah perubahan kode.

### 3.4 Function dan Method pada Struct

AI digunakan untuk memahami perbedaan function biasa dengan method pada struct.

Contoh konsep yang dipelajari:

```go
func (u *User) Activate() {
    u.IsActive = true
}
```

dan:

```go
func (u User) Info() string {
    return ...
}
```

Pembahasan mencakup value receiver dan pointer receiver serta alasan penggunaan pointer receiver ketika data struct perlu diubah.

---

## 4. Cara AI Digunakan dalam Pengerjaan

AI digunakan melalui beberapa bentuk interaksi:

1. **Menanyakan konsep** ketika terdapat materi Go yang belum dipahami.
2. **Meminta penjelasan kode** untuk memahami fungsi setiap bagian.
3. **Menguji pemahaman** dengan memberikan kode yang telah dibuat kemudian menanyakan apakah kode tersebut sudah benar.
4. **Mencari penyebab error** berdasarkan pesan error dan kode yang digunakan.
5. **Membandingkan beberapa cara penulisan kode**, misalnya `var`, `:=`, pointer, dan value parameter.
6. **Meminta contoh sederhana** agar konsep lebih mudah dipahami.
7. **Mempertahankan percobaan kode yang sudah dibuat** dan menggunakan AI untuk menjelaskan fungsi dari setiap percobaan.

---

## 5. Batasan Penggunaan AI

AI digunakan sebagai **alat bantu**, sedangkan kode dan percobaan tetap disesuaikan dengan kebutuhan modul.

Penggunaan AI terutama ditujukan untuk:

* Memahami konsep.
* Menjelaskan sintaks.
* Membantu menemukan kesalahan.
* Membandingkan alternatif implementasi.
* Membantu memahami hubungan antar konsep.

Hasil dari AI tetap diperiksa dan disesuaikan kembali dengan kode yang digunakan dalam project.

---

## 6. Kesimpulan

Penggunaan AI dalam modul Go membantu proses pembelajaran terutama dalam memahami konsep dasar seperti variabel, function, pointer, struct, package, scope, slice, looping, serta penggunaan Go dalam pengembangan backend.

AI digunakan sebagai pendamping dalam proses belajar sehingga setiap percobaan kode dapat dipahami alasan dan cara kerjanya, bukan hanya menghasilkan kode yang dapat dijalankan.
