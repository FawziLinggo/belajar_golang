package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x == 10 && y == 10 {
		fmt.Println("benar")
	} else if x == 5 || y == 6 {
		fmt.Println("Salah")
	}
}
