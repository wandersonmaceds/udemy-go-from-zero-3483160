package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {
	fmt.Println("Functions are awesome!")
	sumResult := sum(10, 20)
	fmt.Println(sumResult)

	// functions are first-citizen type, we can set it to a variable
	// we can also define the variable types as last thing if they all have the same type
	var function = func(n1, n2 int8) int8 {
		return n1 * n2
	}

	functionResult := function(10, 12)
	fmt.Println(functionResult)

	// functions can also have multiple returns.
	var functionWithMultipleReturns = func(n1, n2 int8) (int8, int8) {
		return n1 + n2, n1 * n2
	}

	sumResult, multiplyResult := functionWithMultipleReturns(10, 10)

	fmt.Println(sumResult, multiplyResult)

	// ignore returns by using an underline
	_, multiplyResult2 := functionWithMultipleReturns(10, 10)

	fmt.Println(multiplyResult2)

}
