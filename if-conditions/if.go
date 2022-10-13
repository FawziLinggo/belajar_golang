package main

import "fmt"

func main() {
	name := "fawz"
	if name == "fawzi" {
		fmt.Println(true)
	} else {
		fmt.Println("bukan fawzi")
	}

	if panjangnama := len(name); panjangnama < 5 {
		fmt.Println("nama terlalu panjang")
	}
}
