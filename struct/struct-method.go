package main

import "fmt"

type Costumer2 struct {
	Name, Address string
	Age           int
}

func (costumer Costumer2) costumerName() {
	fmt.Println("Hallo my name is", costumer.Name)
	fmt.Println(" i lived in ", costumer.Address)
}

func main() {
	pelanggan3 := Costumer2{
		Name:    "Rudy",
		Address: "Jakarta",
		Age:     15,
	}

	pelanggan3.costumerName()
}
