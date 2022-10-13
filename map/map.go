package main

import "fmt"

func main() {
	preson := map[string]string{
		"name":    "fawzi Linggo",
		"address": "Jakarta Barat",
	}

	fmt.Println(preson["name"])
	fmt.Println(preson)

	preson["title"] = "Developer"
	fmt.Println(preson)

}
