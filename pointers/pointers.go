package main

import "fmt"

func main() {
	var v1 int8 = 10
	var v2 int8 = v1

	fmt.Println(v1, v2)

	// changing v1 doesn't change v2, because the values are copied when v1 is assigned to v2
	v1 = 20
	fmt.Println(v1, v2)

	// now, with a v3 pointer poiting to v1, any change to v1 will be pointed by v3.
	var v3 *int8 // use * to create a pointer
	v3 = &v1     // use & to capture the pointer to a value

	fmt.Println(v1, v3)  // pointers will print the memory address
	fmt.Println(v1, *v3) // * forces to print the value referenced by the memory address (ex: 0x1400010600a)

	v1 = 30
	fmt.Println(v1, *v3) // prints the same vales, since v1 changed, v3 shows the same value as v1

}
