package main

import "fmt"

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["year"] = "1970"
	b["model"] = "Mustang Shelby"
	fmt.Println("After Change ")
	fmt.Println(a)
	fmt.Println(b)

}
