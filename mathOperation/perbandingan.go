package main

import "fmt"

func main() {
	x := 7
	y := 8

	if x < y {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	name := "Fawzi"
	name_2 := "fawzi"

	var result bool = name == name_2

	fmt.Println(result)
}
