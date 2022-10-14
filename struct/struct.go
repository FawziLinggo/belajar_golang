package main

import "fmt"

type Costumer struct {
	Name, Address string
	Age           int
}

func main() {

	pelanggan1 := Costumer{}
	pelanggan1.Name = "Fawzi"
	pelanggan1.Address = "jakarta"
	// pelanggan1.Age = 15

	fmt.Println(pelanggan1)

	pelanggan2 := Costumer{
		Name: "Linggo",
		// Address: "Takengon",
		Age: 12,
	}

	fmt.Println(pelanggan2)
}
