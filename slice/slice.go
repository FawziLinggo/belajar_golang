package main

import "fmt"

func main() {

	fmt.Println()

	var bulan = [...]string{
		"januari",
		"Februari", "Maret", "April",
		"Mei", "Juni", "Juli",
		"Agustus", "September", "Oktober",
		"November", "Desember",
	}

	fmt.Println(len(bulan))

	slice1 := bulan[4:7]

	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))
	fmt.Println(slice1)

}
