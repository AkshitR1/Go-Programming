package main

import "fmt"

func mymess() {
	fmt.Println("This is my time to enter")
}

func sentcomp(fname string, age int) {
	fmt.Println("Hello my name is ", fname, "and I am ", age, "years old")
}

func add(x int, y int) int {
	return x + y
}
func sub(x int, y int) int {
	return x - y
}
func div(x int, y int) int {
	return x / y
}
func mult(x int, y int) int {
	return x * y
}
func main() {
	sentcomp("AK", 32)
	sentcomp("RK", 38)
	sentcomp("SUMMA", 42)
	mymess()
	fmt.Println(add(123, 324))
	fmt.Println(sub(2424, 2334))
	fmt.Println(div(13123, 3243))
	fmt.Println(mult(12323, 756))
}
