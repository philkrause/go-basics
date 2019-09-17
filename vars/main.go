package main

import "fmt"

//global vars need long form
var size float32 = 10

func main() {
	const name string = "Phil"
	var age int32 = 33
	var musician bool = true
	test := "test"

	firstname, email := "phil", "phil.com"

	fmt.Println(name, age, test, musician, size, firstname, email)
	fmt.Printf("%T\n", test)
}
