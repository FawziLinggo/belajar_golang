package main

import "fmt"

func main() {
	counter := 0
	incremental := func() {
		fmt.Println("incremental")
		counter++
	}

	fmt.Println(counter)
	incremental()
	fmt.Println(counter)
}
