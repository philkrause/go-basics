package main

import "fmt"

func main() {
	//arrays
	// var fruitArr [2]string

	//assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	//declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:4])
}
