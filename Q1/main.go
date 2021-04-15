package main

import (
	"bufio"
	"fmt"
	"os"
)

type Barang struct {
	Nama  string
	Harga int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Nama toko : ")
	scanner.Scan()
	namaToko := scanner.Text()

	fmt.Print("Tanggal transaksi : ")
	scanner.Scan()
	tanggalTransaksi := scanner.Text()

	fmt.Print("Nama kasir : ")
	scanner.Scan()
	namaKasir := scanner.Text()

	var daftarBelanja []Barang

	lanjut := "lanjut"
	for lanjut == "lanjut" {
		var namaBarang string
		var hargaBarang int

		fmt.Println()
		fmt.Print("Nama barang : ")
		fmt.Scan(&namaBarang)
		fmt.Print("Harga barang : ")
		fmt.Scan(&hargaBarang)

		daftarBelanja = append(daftarBelanja, Barang{
			Nama:  namaBarang,
			Harga: hargaBarang,
		})

		fmt.Print("Ingin lanjut atau keluar? : ")
		fmt.Scan(&lanjut)
		if lanjut == "keluar" {
			break
		}
	}
	total := 0
	fmt.Printf("\n%-3s %s\n", "", namaToko)
	fmt.Printf("Tanggal : %s\n", tanggalTransaksi)
	fmt.Printf("%-10s %s\n", "Nama Kasir :", namaKasir)
	fmt.Printf("==============================\n")
	for _, structBelanja := range daftarBelanja {
		fmt.Printf("%-22s Rp %d\n", structBelanja.Nama, structBelanja.Harga)
		total = total + structBelanja.Harga
	}
	fmt.Printf("\n%-23sRp %d\n", "Total :", total)
}
