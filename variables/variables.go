package main

import "fmt"

func main() {
	// we can define variables defining its type or by infering their type. Use the var keywork
	var variable1 string = "Variable 1"
	// we need to use the defined variable or the compile will trow an error for unused variable
	fmt.Println(variable1)

	// by infering its type with no var keywork used
	variable2 := "Variable 2"
	fmt.Println(variable2)

	// we can also do multiple variable declarations separating them with a comma.
	var variable3, variable4 string = "Variable 3", "Variable 4"
	fmt.Println(variable3, variable4)

	variable5, variable6 := "Variable 5", "Variable 6"
	fmt.Println(variable5, variable6)

	// we also have constant, values that cannot be changed once defined.
	const constant1 string = "Constant 1"
	fmt.Println(constant1)

	// variables inversion values can be easilly achieved by direct attribution
	variable1, variable2 = variable2, variable1
	println(variable1, variable2)
}
