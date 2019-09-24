package main

import "fmt"

func main() {
	ids := []int{33, 34, 25, 36, 47, 47}

	//loop thorugh ids

	// for i, id := range ids {
	// 	fmt.Printf("%d - ID: %d\n", i, id)
	// }

	//no index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	//range with map
	emails := map[string]string{"Bob": "Bob@gmail.com", "Mike": "Mike@gmail.com", "Sharon": "sharong@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name:", k)
	}

}
