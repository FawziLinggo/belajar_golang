package main

import (
	"fmt"
)

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("hi")
		counter++
	}

	for value := 1; value <= 5; value++ {
		fmt.Println("nilai ke-", value)
	}

	names := [...]string{"budi", "anduk"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i, value := range names {
		fmt.Println("index - ", i, "=", value)
	}

	for _, value := range names {
		fmt.Println("values :", value)
	}
}
