package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	ID        int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {

	var args = os.Args
	testing := args[1]
	var printStudent Student

	// fmt.Println(printStudent)

	var students = []Student{
		{ID: 1, nama: "Udin", alamat: "Jalan-jalan kuy", pekerjaan: "Mahasiswa", alasan: "Ingin memulai back-end dev"},
		{ID: 2, nama: "Seno", alamat: "Jalan Luwak", pekerjaan: "Barista", alasan: "Ingin memperbaiki hidup"},
		{ID: 3, nama: "Sekar", alamat: "Jalan Mekar selalu", pekerjaan: "Kasir", alasan: "Mau gaji besar"},
		{ID: 4, nama: "Jaja", alamat: "Jalan Van Gogh", pekerjaan: "Pelukis", alasan: "Mau bisa coding"},
		{ID: 5, nama: "Dinda", alamat: "Jalan Balance", pekerjaan: "Accountant", alasan: "Ingin switch karir"},
	}

	for i, student := range students {
		if convert(testing, i) {
			printStudent = student
				fmt.Println("ID:", printStudent.ID)
				fmt.Println("nama:", printStudent.nama)
				fmt.Println("alamat:", printStudent.alamat)
				fmt.Println("pekerjaan:", printStudent.pekerjaan)
				fmt.Println("alasan:", printStudent.alasan)
				break
		}
	}
}

func convert(strNum string, intNum int) bool {
	v, err := strconv.Atoi(strNum)

	var toInt int

	if err != nil {
		panic("Error dari convert function")
	}

	toInt = v

	return toInt == intNum+1 

}

