package main

import "fmt"

type HasName interface {
	GetName() string
}

func saySuiiii(name HasName) {
	fmt.Println("hello ", name.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	eko := Person{
		Name: "Ronaldo",
	}

	saySuiiii(eko)
}
