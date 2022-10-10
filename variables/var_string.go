package main

import "fmt"

func main() {
	var name string
	name = "Fawzi Linggo"
	fmt.Println(name)

	name = "vanger"
	fmt.Println(name)

	var friendsname = "Budi"
	fmt.Println(friendsname)

	var age = 45
	fmt.Println(age)

	country := "America"
	fmt.Println(country)

	var (
		firstName = "Fawzi"
		lastName  = "Linggo"
	)
	fmt.Println("my name is : ", firstName, lastName)

	const mother_name = "sri"

	fmt.Println(mother_name)

	const (
		father_name  = "suci"
		brother_name = "tidak suci"
	)

	fmt.Println(father_name, brother_name)
}
