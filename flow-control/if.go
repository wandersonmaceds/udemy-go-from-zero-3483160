package main

import "fmt"

func main() {
	fmt.Println("Flow Control: IF")

	n := 10

	if n <= 10 {
		println("number is less then or equal to 10")
	}

	// I can initialize a temporary/closure-scoped variable in the if condition
	if n2 := n + 1; n2 >= 10 {
		println("new temp number is greater than or equal to 10")
	}

	if 10 > 11 {
		println("10 > 11")
	} else if 11 > 10 {
		println("11 > 10")
	}

	if 10 > 11 {
		println("10 > 11")
	} else if 11 > 12 {
		println("11 > 10")
	} else {
		println("no numbers here")
	}
}
