package main

import "fmt"

func main() {
	var x int32 = 10
	var y int32 = 10

	if x < y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", x, y)
	}

	//elseif

	color := "red"

	// if color == "red" {
	// 	fmt.Println("color is red")
	// } else if color == "blue" {
	// 	fmt.Println("color is blue")
	// } else {
	// 	fmt.Println("color is not blue nor red")
	// }

	//switch
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not red or blue")
	}

}
