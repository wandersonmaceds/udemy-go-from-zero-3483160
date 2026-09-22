package main

import "fmt"

func main() {

	// arrays must have its length defined, it's required
	// but we can use [...]{} that acts like a lazy-initialization to define the length
	var arr1 [5]int
	arr1[0] = 1

	fmt.Println(arr1)
	fmt.Println(arr1[0])

	arr2 := [...]int{1, 2, 3, 4}
	fmt.Println(arr2)
	// we can'to arr2[5] = 5, it doesn't work, the length was defined in the initialization

	// slides are like array, but with dynamic length.
	var slc1 []int
	fmt.Println(slc1) // prints []

	slc2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slc2)

	// we can change slices by using some functions, but we're actually creating new slices based on others.
	// append creates a new slice based on another one and adding more items
	slc3 := append(slc2, 7)
	fmt.Println(slc3)
}
