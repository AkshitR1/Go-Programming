package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Enter value for a:")
	fmt.Scan(&a)
	fmt.Println("Enter value for b:")
	fmt.Scan(&b)
	if a > b {
		fmt.Println("thats correct!!!")
	} else {
		fmt.Println("thats wrong!!!")
	}

	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}

}
