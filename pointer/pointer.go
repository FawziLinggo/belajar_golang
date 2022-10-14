package main

import "fmt"

type Location struct {
	Kota, Provinsi, Negara string
}

func main() {

	address1 := Location{
		Kota:     "jakarta",
		Provinsi: "DKI jakarta",
		Negara:   "bogor",
	}
	address2 := &address1
	address2.Kota = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Location{
		Kota:     "Malang",
		Provinsi: "Jawa Timur",
		Negara:   "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)

	addres4 := new(Location)
	addres4.Negara = "Iran"
	fmt.Println(addres4)
}
