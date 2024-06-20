package main

import (
	"fmt"
)

func main() {

	var arr = [4]int{22, 34, 2113, 11}

	arr2 := [4]int{90, 84, 4881, 92}

	fmt.Println(arr[2])
	fmt.Println(arr2[0])

	myslice := []string{"Hello", "Good", "Morning"}
	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(myslice)
	fmt.Println(myslice2)
}
