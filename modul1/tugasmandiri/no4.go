package main

import "fmt"

type Student struct {
	ID int			`json: "id"`
	Name string		`json: "nama"`
	Grade float64	`json: "kelas"`
	IsActive bool	`json: "status"`
}

func GetInfo() string {
	info := fmt.Sprintln(siswa1.ID, siswa1.Name, siswa1.Grade, siswa1.IsActive)
	return info
}

func UpdateGrade(grade float64){
	siswa1.Grade = grade
}

func activate(){
	siswa1.IsActive = true
}

func deactive(){
	siswa1.IsActive = false
}

// buat fungsi getinfo baru dengan parameter namaObject
func GetInfoStudent(siswa Student) string {
	info := fmt.Sprintln(
		siswa.ID,
		siswa.Name,
		siswa.Grade,
		siswa.IsActive,
	)
	return info
}

// deklarasi object dalam slice untuk looping
var daftarSiswa = []Student{
	siswa1,
	siswa2,
	siswa3,
}

// Menampilkan semua student dengan looping
func GetAllStudent() {
	for i, siswa := range daftarSiswa {
		fmt.Println("Student ke-", i+1)
		fmt.Println("ID:", siswa.ID)
		fmt.Println("Nama:", siswa.Name)
		fmt.Println("Grade:", siswa.Grade)
		fmt.Println("Status:", siswa.IsActive)
		fmt.Println()
	}
}

// Mengubah grade object Student tertentu
func UpdateGradeStudent(siswa *Student, grade float64) {
	siswa.Grade = grade
}

// Mengaktifkan student tertentu
func ActivateStudent(siswa *Student) {
	siswa.IsActive = true
}

// Menonaktifkan student tertentu
func DeactiveStudent(siswa *Student) {
	siswa.IsActive = false
}
