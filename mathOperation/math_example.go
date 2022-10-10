package main

import "fmt"

func main() {
	var (
		apple  = 13
		orange = 29
		total  = apple + orange
		result = 10 + 10
	)

	fmt.Println(result)
	fmt.Println(total)

	result *= result
	fmt.Println("result * ", result)
}
