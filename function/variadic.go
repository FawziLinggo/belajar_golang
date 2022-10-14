package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := sumAll(10, 10, 10, 10)
	fmt.Println(result)

	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(sumAll(slice...))
}
