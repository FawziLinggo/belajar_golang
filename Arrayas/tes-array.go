package main

import "fmt"

func main() {

	var (
		name [3]string
	)

	name[0] = "fawzi"
	name[1] = "Linggo"
	fmt.Println(name[0])

	var values = [3]int{
		40,
		12,
		17,
	}
	fmt.Println(len(values))
}
