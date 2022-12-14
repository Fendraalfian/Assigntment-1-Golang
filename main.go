package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	absen		    int
	nama                string
	alamat              string
	pekerjaan           string
	alasanMemilihGolang string
}

func main() {

	argument := os.Args
	printData(argument)

}

func printData(arguments []string) {

	startIndex, _ := strconv.Atoi(arguments[1])

	for _, value := range dummyData() {
		if value.absen == startIndex {
			fmt.Println("=====Data Siswa====")
			fmt.Println("Nama:", value.nama)
			fmt.Println("Alamat:", value.alamat)
			fmt.Println("Pekerjaan:", value.pekerjaan)
			fmt.Println("Alasan memilih golang:", value.alasanMemilihGolang)
			fmt.Println("===================")
			return
		}
	}

	fmt.Println("Siswa tidak ditemukan")

}

func dummyData() []student {
	students := []student{
		{1, "Fendra Alfian Widiyanto", "Manado", "Freshgraduate", "Karena bahasa yang populer dan diminati banyak perusahaan"},
		{2, "Devin Booker", "US", "Software Engineer", "Karena bahasa yang populer dan diminati banyak perusahaan"},
		{3, "Giannis Antetokounmpo", "Greece", "UX Design", "Bahasa yang sangat wajib dipelajari karena sangat populer"},
		{4, "Luka Doncic", "Slovenia", "Android Developer", "Banyak perusahaan yang mencari tenaga ahli pada bahasa golang"},
		{5, "Luka Jokic", "Serbia", "Dev Ops", "Karena bahasa yang populer dan diminati banyak perusahaan"},
	}

	return students
}
