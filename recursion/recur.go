package main

import "fmt"

func fact(x float64) (y float64) {
	if x > 0 {
		y = x * fact(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	fmt.Println(fact(1))
	fmt.Println(fact(2))
	fmt.Println(fact(3))
	fmt.Println(fact(4))
	fmt.Println(fact(5))
	fmt.Println(fact(6))
	fmt.Println(fact(7))
	fmt.Println(fact(8))
	fmt.Println(fact(9))

}
