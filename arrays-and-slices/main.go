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

	// slices internally works like Java Lists
	// when can reach the initial length limit (internal array),
	// a new bigger array is created and all values are copied
	fmt.Println("Slice Internals")
	var slice []int8
	fmt.Println(slice, len(slice), cap(slice)) // printing length and capacity, starts with 0 and 0

	slice = append(slice, 1)                   // adding an element to the slice
	fmt.Println(slice, len(slice), cap(slice)) // printing length and capacity, now we have 1 and 8

	fmt.Println("Slice Internals: Make")
	// we can also create slices with pre-defined lengh and capacity with the make function
	slice2 := make([]int8, 2, 3) // length 2, capacity 3
	fmt.Println(slice2, len(slice2), cap(slice2))
	slice2 = append(slice2, 1)
	slice2 = append(slice2, 1)
	fmt.Println(slice2, len(slice2), cap(slice2)) // length overflows the initial 2 to 4 and capacity goes up to 8
}
