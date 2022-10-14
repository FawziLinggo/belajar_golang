package main

import "fmt"

type Location2 struct {
	Kota, Provinsi, Negara string
}

func changeNegaraToIndonesia(negara *Location2) {
	negara.Negara = "Indonesia"
}

func main() {
	address := Location2{
		Kota:     "jakarta",
		Provinsi: "DKI jakarta",
		Negara:   "bogor",
	}

	// Tidak Berubah
	// changeNegaraToIndonesia(address)
	// fmt.Println(address)

	changeNegaraToIndonesia(&address)
	fmt.Println(address)

}
