package main

import "fmt"

// func main() {
// 	var day int
// 	fmt.Println("Enter the day of week(1 to n):")
// 	fmt.Scan(&day)
// 	switch day {
// 	case 1:
// 		fmt.Println("Monday")
// 	case 2:
// 		fmt.Println("Tuesday")
// 	case 3:
// 		fmt.Println("Wednesday")
// 	case 4:
// 		fmt.Println("Thursday")
// 	case 5:
// 		fmt.Println("Friday")
// 	case 6:
// 		fmt.Println("Saturday")
// 	case 7:
// 		fmt.Println("Sunday")
// 	default:
// 		fmt.Println("Not a weekday")
// 	}
// }

func main() {
	var day int
	fmt.Println("Enter the day of week(1 to 7):")
	fmt.Scan(&day)
	switch day {
	case 1, 3, 5:
		fmt.Println("Odd Weekday")
	case 2, 4:
		fmt.Println("Even Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Enter valid number")
	}
}
