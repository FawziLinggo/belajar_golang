package main

import (
	"fmt"
	"main/databases"
)

func main() {

	result := databases.GetDatabase()
	fmt.Println(result)
}
