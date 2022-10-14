package main

import "fmt"

func logging() {
	fmt.Println("Ini adalah loging (defer). app running")
	message := recover()
	if message != nil {
		fmt.Println("terjadi  error dengan pesan : ", message)
	}
}
func runApp(error bool) {
	defer logging()
	fmt.Println("run app")
	if error {
		panic("Yamette kude..")
	}

}

func main() {

	runApp(false)

}
