package main

import (
	"strconv"
	"strings"
	"sort"

	"github.com/gofiber/fiber/v2"
)

var students []Student
var nextID = 1

// requireJSON menolak request berisi body yang Content-Type-nya bukan JSON.
// Status yang tepat untuk kasus ini adalah 415, bukan 400.
func requireJSON(c *fiber.Ctx) error {
    if metodeBerbody[c.Method()] {
        ct := c.Get("Content-Type")
        if !strings.HasPrefix(ct, fiber.MIMEApplicationJSON) {
            return fail(c, fiber.StatusUnsupportedMediaType,
                "Content-Type harus application/json")
        }
    }
    return c.Next()
}

func findStudentIndex(id int) int {
	for i := range students {
		if students[i].ID == id {
			return i
		}
	}
	return -1 //eror untuk data tidak ditemukan
}

// mendeteksi NIM ganda
func findStudentByNIM(nim int) int {
	for i := range students {
		if students[i].NIM == nim {
			return i
		}
	}
	return -1
}

func paramID(c *fiber.Ctx) (int, bool) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id < 1 {
		return 0, false
	}
	return id, true
}

func getAllStudents(c *fiber.Ctx) error {
	// mengambil page dan limit dari query string
	// default page = 1, limit = 10
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	// page dan limit harus positif
	if page < 1 {
		return fail(c, fiber.StatusBadRequest, "page harus berupa angka positif")
	}

	if limit < 1 {
		return fail(c, fiber.StatusBadRequest, "limit harus berupa angka positif")
	}

	// batas maksimal limit = 100
	// agar client tidak meminta terlalu banyak data dalam satu request
	if limit > 100 {
		return fail(c, fiber.StatusBadRequest, "limit maksimal 100")
	}

	// mengambil seluruh data mahasiswa
	data := students

	// search berdasarkan nama, tidak membedakan huruf besar dan kecil
	search := strings.TrimSpace(c.Query("search"))
	if search != "" {
		var filtered []Student

		for _, student := range data {
			if strings.Contains(
				strings.ToLower(student.Name),
				strings.ToLower(search),
			) {
				filtered = append(filtered, student)
			}
		}

		data = filtered
	}

	// filter berdasarkan is_active
	isActive := c.Query("is_active")
	if isActive != "" {
		var filtered []Student

		for _, student := range data {
			if isActive == "true" && student.IsActive {
				filtered = append(filtered, student)
			} else if isActive == "false" && !student.IsActive {
				filtered = append(filtered, student)
			}
		}

		if isActive != "true" && isActive != "false" {
			return fail(c, fiber.StatusBadRequest,
				"is_active harus true atau false")
		}

		data = filtered
	}

	// mengambil parameter sort dan order
	sortField := c.Query("sort", "id")
	order := strings.ToLower(c.Query("order", "asc"))

	// whitelist field yang boleh digunakan untuk sorting
	allowedSort := map[string]bool{
		"id":        true,
		"name":      true,
		"grade":     true,
		"is_active": true,
		"nim":       true,
	}

	if !allowedSort[sortField] {
		return fail(c, fiber.StatusBadRequest,
			"field sort tidak diizinkan")
	}

	if order != "asc" && order != "desc" {
		return fail(c, fiber.StatusBadRequest,
			"order harus asc atau desc")
	}

	// sorting berdasarkan field yang sudah masuk whitelist
	sort.Slice(data, func(i, j int) bool {
		var hasil bool

		switch sortField {
		case "id":
			hasil = data[i].ID < data[j].ID
		case "name":
			hasil = strings.ToLower(data[i].Name) < strings.ToLower(data[j].Name)
		case "grade":
			hasil = data[i].Grade < data[j].Grade
		case "is_active":
			hasil = !data[i].IsActive && data[j].IsActive
		case "nim":
			hasil = data[i].NIM < data[j].NIM
		}

		if order == "desc" {
			return !hasil
		}

		return hasil
	})

	// jumlah data setelah search dan filter
	total := len(data)

	// menghitung jumlah halaman
	totalPages := (total + limit - 1) / limit

	if totalPages == 0 {
		totalPages = 1
	}

	if page > totalPages {
		return fail(c, fiber.StatusBadRequest,
			"page melebihi jumlah halaman")
	}

	// menentukan data sesuai page dan limit
	start := (page - 1) * limit
	end := start + limit

	if end > total {
		end = total
	}

	data = data[start:end]

	// informasi pagination
	meta := &Meta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}

	return okList(c, "daftar mahasiswa berhasil diambil", data, meta)
}

func getStudent(c *fiber.Ctx) error {
	id, valid := paramID(c)
	if !valid {
		return fail(c, fiber.StatusBadRequest, "id harus berupa angka positif")
	}

	i := findStudentIndex(id)
	if i == -1 {
		return fail(c, fiber.StatusNotFound, "mahasiswa tidak ditemukan")
	}

	return ok(c, "mahasiswa ditemukan", students[i])
}

func postCreatetStudent(c *fiber.Ctx) error {
	var req requestPostCreateStudent
	if err := c.BodyParser(&req); err != nil {
		return fail(c, fiber.StatusBadRequest, "body harus berupa JSON yang valid")
	}

	errs := map[string]string{}
	req.Name = strings.TrimSpace(req.Name)

	if req.Name == "" {
		errs["name"] = "wajib diisi"
	}
	if req.Grade < 0 || req.Grade > 100 {
		errs["grade"] = "harus antara 0 sampai 100"
	}
	if req.NIM < 1 {
		errs["nim"] = "wajib diisi dan berupa angka positif"
	}
	if len(errs) > 0 {
		return failValidation(c, errs)
	}

	// NIM ganda = bertentangan dengan data yang ada (409), dicek setelah validasi lolos
	if findStudentByNIM(req.NIM) != -1 {
		return fail(c, fiber.StatusConflict, "NIM sudah terdaftar")
	}

	baru := Student{
		ID:       nextID,
		Name:     req.Name,
		Grade:    req.Grade,
		IsActive: req.IsActive,
		NIM:      req.NIM,
	}
	students = append(students, baru)
	nextID++

	return created(c, "mahasiswa berhasil ditambahkan", baru,
		"/api/v1/students/"+strconv.Itoa(baru.ID))
}


// PUT mengganti SELURUH isi. Field yang tidak dikirim dianggap dikosongkan,
// karena itu semuanya wajib ada. --> kolom yg gaada inputnya akan dikosongkan atau diisi nilai default
func putEditSeluruh(c *fiber.Ctx) error {
	id, valid := paramID(c)
	if !valid {
		return fail(c, fiber.StatusBadRequest, "id harus berupa angka positif")
	}

	i := findStudentIndex(id)
	if i == -1 {
		return fail(c, fiber.StatusNotFound, "mahasiswa tidak ditemukan")
	}

	var req requestPutEditSeluruhStudent
	if err := c.BodyParser(&req); err != nil {
		return fail(c, fiber.StatusBadRequest, "body harus berupa JSON yang valid")
	}

	// bagian yang menyatakan kolom yg harus diisi
	errs := map[string]string{}
	//jika input kolom name tidak ada maka akan berisi default zero value string yaitu ""
	if strings.TrimSpace(req.Name) == "" {
		// pesan error saat kolom name zero value :
		errs["name"] = "wajib diisi pada PUT"
	}
	if req.Grade < 0 || req.Grade > 100 {
		errs["grade"] = "harus antara 0 sampai 100"
	}
	if len(errs) > 0 {
		return failValidation(c, errs)
	}

	students[i].Name = strings.TrimSpace(req.Name)
	students[i].Grade = req.Grade
	// jika kolom isActive yg bertipe data boolean kosong maka akan akan diisi zero value yaitu false
	students[i].IsActive = req.IsActive

	return ok(c, "mahasiswa berhasil diganti seluruhnya", students[i])
}

// PATCH hanya mengubah field yang benar-benar dikirim.
func patchEditSebagian(c *fiber.Ctx) error {
	id, valid := paramID(c)
	if !valid {
		return fail(c, fiber.StatusBadRequest, "id harus berupa angka positif")
	}

	i := findStudentIndex(id)
	if i == -1 {
		return fail(c, fiber.StatusNotFound, "mahasiswa tidak ditemukan")
	}

	var req requestPatchEditSebagianStudent
	if err := c.BodyParser(&req); err != nil {
		return fail(c, fiber.StatusBadRequest, "body harus berupa JSON yang valid")
	}

	// Grade bertipe int biasa, jadi dianggap tidak dikirim kalau 0
	if req.Name == nil && req.IsActive == nil && req.Grade == 0 {
		return fail(c, fiber.StatusBadRequest, "tidak ada field yang diubah")
	}

	if req.Name != nil {
		if strings.TrimSpace(*req.Name) == "" {
			return failValidation(c, map[string]string{"name": "tidak boleh kosong"})
		}
		students[i].Name = strings.TrimSpace(*req.Name)
	}
	if req.Grade < 0 || req.Grade > 100 {
		return failValidation(c, map[string]string{"grade": "harus antara 0 sampai 100"})
	}
	students[i].Grade = req.Grade // harus diisi, kalau kosong dioverwrite 0 (sesuai rancanganmu)
	if req.IsActive != nil {
		students[i].IsActive = *req.IsActive
	}

	return ok(c, "mahasiswa berhasil diperbarui sebagian", students[i])
}

func deleteStudent(c *fiber.Ctx) error {
	id, valid := paramID(c)
	if !valid {
		return fail(c, fiber.StatusBadRequest, "id harus berupa angka positif")
	}

	i := findStudentIndex(id)
	if i == -1 {
		return fail(c, fiber.StatusNotFound, "mahasiswa tidak ditemukan")
	}

	students = append(students[:i], students[i+1:]...)

	return noContent(c) // 204: berhasil, dan memang tidak ada yang perlu dikirim
}
