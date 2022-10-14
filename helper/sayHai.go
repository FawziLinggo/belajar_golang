package helper

import "fmt"

func SayHai(name string) {
	fmt.Println("Hai", name)
}

// tidak bisa di akses dari luar
func sayGoodBye(name string) {
	fmt.Println("good bye ", name)
}
