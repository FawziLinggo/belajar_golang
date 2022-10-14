package main

import "fmt"

type (
	Filtering func(string) string
)

func sayHello(name string, filter Filtering) {
	nameFiltered := filter(name)
	fmt.Println(" Hello ", nameFiltered)
}

func sayHelloFiltering(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHello("Anjing", sayHelloFiltering)
}
