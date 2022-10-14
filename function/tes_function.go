package main

import "fmt"

func sayHello() {
	fmt.Println("hello")
}

func hello(firstname string, lastname string) {
	fmt.Println("Hello ", firstname+lastname)
}

func getName(name string) string {
	return name
}

func getFullname() (string, string) {
	return "fawzi", "Linggo"
}

func getComplateName() (firstName int, lastName string) {
	firstName = 123
	lastName = "Linggo"

	return
}

func main() {
	sayHello()
	result := getName("kuniatod")
	hello("Fawzi", "Linggo")
	fmt.Println(result)

	// fistname, lastname := getFullname()
	// fmt.Println(fistname, lastname)

	fistname, _ := getFullname()
	fmt.Println(fistname)

	a, b := getComplateName()
	fmt.Println("complete :", a, b)
}
