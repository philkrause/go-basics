package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", a)

	//use to read value from address
	fmt.Println(*b)
	fmt.Println(*&a)

	//change val with value
	*b = 10
	fmt.Println(a)
}
