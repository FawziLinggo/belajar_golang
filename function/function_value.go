package main

import "fmt"

func sayGodbay(name string) string {
	return "good bye " + name
}

func main() {
	goodbye := sayGodbay
	fmt.Println(goodbye("fawzi"))
}
