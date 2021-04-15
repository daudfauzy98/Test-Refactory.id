package main

import "fmt"

func isLeap(year int) bool {
	// Jika value dari variabel "year" memenuhi ketentuan dibawah,
	// secara deafult akan mengirimkan nilai true
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}

func main() {
	var tahunPertama, tahunKedua int

	fmt.Print("Masukkan tahun pertama : ")
	fmt.Scan(&tahunPertama)

	fmt.Print("Masukkan tahun kedua : ")
	fmt.Scan(&tahunKedua)
	fmt.Println("\n")

	delimiter := ""
	for i := tahunPertama; i <= tahunKedua; i++ {
		if isLeap(i) {
			fmt.Printf("%s%d", delimiter, i)
			delimiter = ", "
		}
	}
}
