package main

import (
	"fmt"
	"math"
)

func main() {
	// float
	x := 45
	y := 20
	z := x - y
	fmt.Println("Square root of X-Y is:", math.Sqrt(float64(z)))

	var message string
	// string
	message = "This is a demo program"
	fmt.Println("The message is:", message)

	var boolt bool = true
	// boolean
	fmt.Printf("The answers is %t\n", boolt)

	var age, id int
	// integer
	age = 25
	id = 1
	fmt.Printf("ID is %d , Age is %d \n ", id, age)
}
