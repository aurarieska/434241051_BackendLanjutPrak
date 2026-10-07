# AI Usage Log --- Project Go

Dokumentasi singkat penggunaan AI selama pengerjaan dan pembelajaran project Go.

AI digunakan terutama sebagai pendamping belajar, penjelasan konsep, debugging, perbaikan kode, dan membantu memahami implementasi API. Kode dan struktur project tetap disesuaikan dengan kebutuhan tugas/project. Juga digunakan untuk menyusun penjelasan dalam laporan karena saya sudah menyusun penjelasan  dalam komentar code berdasarkan pemahaman saya, hal ini dilakukan untuk mempercepat hanya perubahan format penjelasan agar waktu dapat digunakan semaksimal mungkin memahami materi. 

## 1. Awal Pembelajaran Go & Struct

- Memahami dasar Go: `package main`, `import`, variabel, `var`, `:=`, zero value, slice, map, pointer, goroutine, dan struct.
- Membahas struct `User` dan kemudian struct `Student`.
- Memahami perbedaan struct type dengan struct value/object.
- Membahas value receiver dan pointer receiver.
- Membahas JSON tag pada struct.

Contoh struct yang digunakan:
```go
type Student struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Grade    int    `json:"grade"`
    IsActive bool   `json:"isactive"`
    NIM      int    `json:"nim"`
}
```

## 2. Belajar Fiber & Project Backend

- Menggunakan Go Fiber v2 untuk membuat REST API.
- Memahami `fiber.New()`, route, `*fiber.Ctx`, response, status code, dan `app.Listen(":3000")`.
- Memahami struktur project/module Go dan penggunaan `go.mod` serta `go.sum`.
- Membahas endpoint API dan pengelompokan route dengan:
  - `app.Group("/api/v1")`
  - `api.Group("/students")`
- Memahami bahwa route group digunakan untuk membuat prefix URL sehingga route di dalamnya tidak perlu menuliskan prefix secara berulang.

## 3. API Student

Mengerjakan API Student dengan operasi:
- `GET /students` --- mengambil seluruh student.
- `GET /students/:id` --- mengambil student berdasarkan ID.
- `POST /students` --- membuat student.
- `PUT /students/:id` --- mengganti seluruh data student.
- `PATCH /students/:id` --- mengubah sebagian data student.
- `DELETE /students/:id` --- menghapus student.

Data sementara disimpan menggunakan slice `[]Student` dan ID dibuat dengan `nextID`.

## 4. Request, Validation, dan Response

- Membahas struct request seperti `CreateStudentRequest` dan `PatchStudentRequest`.
- Memahami penggunaan pointer (`*`) pada field PATCH untuk membedakan field yang tidak dikirim (`nil`) dengan zero value seperti `0` atau `false`.
- Membahas `paramID()` untuk membaca dan memvalidasi parameter `:id`.
- Membahas `requireJSON()` untuk menolak request dengan `Content-Type` yang bukan JSON menggunakan status `415`.
- Membahas helper response seperti:
  - `ok()`
  - `okList()`
  - `created()`
  - `noContent()`
  - `fail()`
- Memahami fungsi `fail()` sebagai helper untuk mengirim error dengan status dan message yang diberikan melalui parameter.

### WebResponse dan Meta

- Membahas `WebResponse` sebagai struktur standar untuk membungkus response API, yang terdiri dari `Success`, `Message`, `Data`, `Meta`, dan `Errors`.
- Membahas `Meta` sebagai informasi tambahan mengenai data, terutama untuk pagination, seperti `page`, `limit`, `total`, dan `total_pages`.
- Memahami bahwa `Meta` menggunakan pointer (`*Meta`) agar dapat bernilai `nil` ketika response tidak membutuhkan metadata.
- Memahami hubungan `*Meta` dengan `omitempty`: jika `Meta` bernilai `nil`, field `meta` tidak dikirim dalam JSON.
- Memahami bahwa `nil` pada `*Meta` bukan berarti mempertahankan nilai lama, karena `WebResponse` merupakan response baru.
- Membahas `okList()` sebagai helper untuk response list yang juga mengirim `Meta`.
- Membahas `failValidation()` untuk mengirim error validasi dengan status `422 Unprocessable Entity` dan detail error melalui field `Errors`.
- Memahami parameter `c` dengan tipe `*fiber.Ctx` sebagai context dari request Fiber yang diteruskan dari handler ke helper response agar helper dapat mengatur dan mengirim response kepada client.
- Memahami `fiber.StatusUnsupportedMediaType` sebagai konstanta untuk HTTP `415 Unsupported Media Type`.

## 5. PUT vs PATCH

- PUT digunakan untuk mengganti keseluruhan data.
- PATCH digunakan untuk mengubah sebagian data.
- Dibahas kembali perbedaan field biasa dan pointer pada request, terutama untuk mengetahui apakah nilai dikirim atau tidak.

## 6. Pagination GET All

- Requirement GET seluruh student dibahas kembali.
- Pagination ditambahkan pada endpoint GET all students.
- Dibahas penggunaan query seperti `page` dan `limit`, serta metadata pagination pada response.

## 7. NIM Unique

- Dibahas bahwa `NIM` tetap berupa field biasa pada struct:
```go
NIM int `json:"nim"`
```

- Karena data masih disimpan di slice, keunikan NIM perlu dicek dalam logic saat membuat student sebelum `append`.
- Jika menggunakan database nantinya, NIM juga sebaiknya diberi constraint `UNIQUE`.

## 8. Routing dan Alur Endpoint

- Membahas `api := app.Group("/api/v1")` sebagai prefix utama API.
- Membahas `s := api.Group("/students", requireJSON)` sebagai grup khusus resource student yang menambahkan prefix `/students` sekaligus menerapkan middleware `requireJSON` pada endpoint di dalam grup tersebut.
- Memahami bahwa kedua prefix tersebut digabungkan, sehingga route seperti `s.Get("/", getAllStudents)` menjadi `GET /api/v1/students`.
- Membahas seluruh routing CRUD sebagai satu kesatuan:
  - `s.Get("/", getAllStudents)`
  - `s.Get("/:id", getStudent)`
  - `s.Post("/", postCreatetStudent)`
  - `s.Put("/:id", putEditSeluruh)`
  - `s.Patch("/:id", patchEditSebagian)`
  - `s.Delete("/:id", deleteStudent)`
- Memahami alur request: client mengirim metode HTTP dan URL → Fiber mencocokkan route → middleware `requireJSON` dijalankan → request diteruskan ke handler yang sesuai.
- Memahami bahwa `/:id` merupakan parameter route, sehingga URL seperti `/api/v1/students/1` akan diteruskan ke handler dengan nilai `id = 1`.
- Membahas endpoint `/api/v1/health` sebagai endpoint tambahan untuk mengecek apakah server berjalan, bukan bagian dari requirement CRUD Student.
- Membahas endpoint `/` dan middleware global seperti request ID, logger, dan CORS sebagai bagian tambahan dari konfigurasi aplikasi, bukan implementasi utama CRUD Student.

