package main

import "fmt"

func factorialRecrusive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecrusive(value-1)
	}
}

func main() {
	result := factorialRecrusive(5)
	fmt.Println(result)
}
