## Dokumen Kontrak API

### 1. Health Check

- **Metode:** `GET`
- **Endpoint:** `/api/v1/health`
- **Parameter:** Tidak ada
- **Body:** Tidak ada
- **Status:**
  - `200 OK`
- **Contoh Response:**

```json
{
  "success": true,
  "message": "server berjalan",
  "data": {
    "timestamp": "2026-10-08T..."
  }
}
```

---

### 2. Get All Students

- **Metode:** `GET`
- **Endpoint:** `/api/v1/students/`
- **Parameter Query:**
  - `page` → nomor halaman, default `1`
  - `limit` → jumlah data per halaman, default `10`, maksimal `100`
  - `search` → pencarian berdasarkan nama
  - `is_active` → filter `true` atau `false`
  - `sort` → `id`, `name`, `grade`, `is_active`, atau `nim`
  - `order` → `asc` atau `desc`
- **Body:** Tidak ada
- **Status:**
  - `200 OK`
  - `400 Bad Request`
- **Contoh Request:**

```text
GET /api/v1/students/?page=1&limit=10&search=Budi&is_active=true&sort=name&order=asc
```

- **Contoh Response:**

```json
{
  "success": true,
  "message": "daftar mahasiswa berhasil diambil",
  "data": [
    {
      "id": 1,
      "name": "Budi",
      "grade": 85,
      "isactive": true,
      "nim": 12345
    }
  ],
  "meta": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "total_pages": 1
  }
}
```

---

### 3. Get Student by ID

- **Metode:** `GET`
- **Endpoint:** `/api/v1/students/:id`
- **Parameter:**
  - `id` → ID mahasiswa
- **Body:** Tidak ada
- **Status:**
  - `200 OK`
  - `400 Bad Request`
  - `404 Not Found`
- **Contoh Request:**

```text
GET /api/v1/students/1
```

- **Contoh Response:**

```json
{
  "success": true,
  "message": "mahasiswa ditemukan",
  "data": {
    "id": 1,
    "name": "Budi",
    "grade": 85,
    "isactive": true,
    "nim": 12345
  }
}
```

---

### 4. Create Student

- **Metode:** `POST`
- **Endpoint:** `/api/v1/students/`
- **Parameter:** Tidak ada
- **Body:** JSON
- **Status:**
  - `201 Created`
  - `400 Bad Request`
  - `409 Conflict`
  - `415 Unsupported Media Type`
  - `422 Unprocessable Entity`
- **Contoh Body:**

```json
{
  "name": "Budi",
  "grade": 85,
  "isactive": true,
  "nim": 12345
}
```

- **Contoh Response:**

```json
{
  "success": true,
  "message": "mahasiswa berhasil ditambahkan",
  "data": {
    "id": 1,
    "name": "Budi",
    "grade": 85,
    "isactive": true,
    "nim": 12345
  }
}
```

---

### 5. Update Student dengan PUT

- **Metode:** `PUT`
- **Endpoint:** `/api/v1/students/:id`
- **Parameter:**
  - `id` → ID mahasiswa
- **Body:** JSON
- **Status:**
  - `200 OK`
  - `400 Bad Request`
  - `404 Not Found`
  - `415 Unsupported Media Type`
  - `422 Unprocessable Entity`
- **Contoh Body:**

```json
{
  "name": "Budi",
  "grade": 90,
  "isactive": true
}
```

- **Contoh Response:**

```json
{
  "success": true,
  "message": "mahasiswa berhasil diganti seluruhnya",
  "data": {
    "id": 1,
    "name": "Budi",
    "grade": 90,
    "isactive": true,
    "nim": 12345
  }
}
```

---

### 6. Update Student dengan PATCH

- **Metode:** `PATCH`
- **Endpoint:** `/api/v1/students/:id`
- **Parameter:**
  - `id` → ID mahasiswa
- **Body:** JSON
- **Status:**
  - `200 OK`
  - `400 Bad Request`
  - `404 Not Found`
  - `415 Unsupported Media Type`
  - `422 Unprocessable Entity`
- **Contoh Body:**

```json
{
  "name": "Budi Baru",
  "grade": 95
}
```

- **Contoh Response:**

```json
{
  "success": true,
  "message": "mahasiswa berhasil diperbarui sebagian",
  "data": {
    "id": 1,
    "name": "Budi Baru",
    "grade": 95,
    "isactive": true,
    "nim": 12345
  }
}
```

---

### 7. Delete Student

- **Metode:** `DELETE`
- **Endpoint:** `/api/v1/students/:id`
- **Parameter:**
  - `id` → ID mahasiswa
- **Body:** Tidak ada
- **Status:**
  - `204 No Content`
  - `400 Bad Request`
  - `404 Not Found`
- **Contoh Request:**

```text
DELETE /api/v1/students/1
```

- **Contoh Response:**

```text
Tidak ada body response.
```

---

## Contoh Response Error

### 400 Bad Request

```json
{
  "success": false,
  "message": "id harus berupa angka positif"
}
```

### 404 Not Found

```json
{
  "success": false,
  "message": "mahasiswa tidak ditemukan"
}
```

### 409 Conflict

```json
{
  "success": false,
  "message": "NIM sudah terdaftar"
}
```

### 415 Unsupported Media Type

```json
{
  "success": false,
  "message": "Content-Type harus application/json"
}
```

### 422 Unprocessable Entity

```json
{
  "success": false,
  "message": "validasi gagal",
  "errors": {
    "name": "wajib diisi",
    "grade": "harus antara 0 sampai 100"
  }
}
```