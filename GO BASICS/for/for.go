package main

import "fmt"

func main() {

	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < i; j++ {

	// 		fmt.Print("*")
	// 	}
	// 	fmt.Print("\n")
	// }
	for i := 0; i < 6; i++ {
		for j := 0; j < 2*i; j++ {
			fmt.Print("")
		}
		for k := 0; k < 2*(4-i)-1; k++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

	// for i := 0; i < 5; i++ {
	// 	if i == 3 {
	// 		fmt.Println("Hello World")
	// 	}
	// 	fmt.Println(i)
	// }
}
