package main

/*
NO. 1c
buat struct student
*/
type Student struct{
	//nama di json dan code boleh beda, tp saya mau samain aja biar gak pusing makenya
	ID			int 	`json:"id"`
	Name 		string	`json:"name"`
	Grade 		int 	`json:"grade"`
	IsActive	bool	`json:"isactive"`
	NIM			int 	`json:"nim"`
}

/*
NO. 1e
struct request post, put, patch
*/
type requestPostCreateStudent struct{ // asumsi nim bikin sendiri, yg generate id aja trs nanti dibikin contstraint unique
	Name 		string	`json:"name"`
	Grade 		int 	`json:"grade"`
	IsActive	bool	`json:"isactive"`
	NIM			int 	`json:"nim"`
}

type requestPutEditSeluruhStudent struct{ // nim gaboleh diedit, setiap kolom klo kosong akan dioverwrite zero value
	Name 		string	`json:"name"`
	Grade 		int 	`json:"grade"`
	IsActive	bool	`json:"isactive"`
}

type requestPatchEditSebagianStudent struct{ // nim gaboleh diedit 
	Name 		*string	`json:"name"` // status boleh gadiisi dan mempertahankan nilai aslinya.
	Grade 		int 	`json:"grade"` // harus diisi, klo kosonng akan dioverwrite 0
	IsActive	*bool	`json:"isactive"` //status boleh gadiisi dan mempertahankan nilai aslinya.
}
/* sebelum dipindah ke laporan saya take notes di code dan gak saya hapus supaya lebih gampang belajarnya gak buka banyak file.
notes perbedaan besar :
PUT vs PATCH : 
PUT : mengubah semua variabel, klo kosong (nil) akan tetap dikirim dan dioverwrite dengan zero value
PATCH : mengubah sebagian variabel, klo kosong (nil) tidak dikirim (membiarkan kosong atau kalau ada nilai lama maka mempertahankan nilai lama )
TAPI TAPI TAPI
pada patch type data bisa diawali * (misal : *int)--> fungsinya : 
Struct	Request	Yang diterima struct		Hasil
Age 	int				{}				0 Age jadi 0
Age 	*int			{}				nil	Age tetap 20
Age 	int			{"age": 0}			0	Age jadi 0
Age 	*int		{"age": 0}			0	Age jadi 0
intinya, 
tanpa * : kosong dianggap 0 --> dikirim zero value
dengan * : kosong dianggap nil --> tidak dikirim

hmm sebenernya yg bikin susah bedain put dan patch example codenya,
disitu put dan patch memang beda tapi validasinya membuat mirip cara kerjanya
put dan patch saja : put kosong akan dioverwrite dengan zero value dan patch kosong tidak dikirim (membiarkan kosong atau kalau ada nilai lama maka mempertahankan nilai lama )
validasi :))))     : put "" --> menolak kirim, tidak dikirim jadi yaa sama aja "" cara kerjanya kaya nil di patch yaitu tidak dikirim. dan patch nil dikasih pesan peringartan.
SOOOOOOOO, mending bedain lewat validasi aja drpd pusing bedain patch dan put, lebih customable. 
*/

/*
NO. 1e
buat amplop response
*/
type Meta struct {
    Page       int `json:"page"`
    Limit      int `json:"limit"`
    Total      int `json:"total"`
    TotalPages int `json:"total_pages"`
}

type WebResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Data    any    `json:"data,omitempty"`
    Meta    *Meta  `json:"meta,omitempty"`
    Errors  any    `json:"errors,omitempty"`
}



