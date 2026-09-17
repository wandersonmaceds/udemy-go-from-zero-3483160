package main

import (
	"errors"
	"fmt"
)

func main() {
	// first we have strings, just a chain of characters:
	var stringType string = "My first string"
	fmt.Println(stringType)

	// we don't have chars like Java, chars here are single quoted
	// they are transformed to integers from the ASCII table.
	char := 'A'       // 65 position on ASCII table
	fmt.Println(char) // will print 65

	// in the numbers side, we have ints and floats.
	// Int can have a length restriction: int8, int16, int32 and int64
	// We can also have unassined ints, uint8 uint16, uint32 and uint64
	// In the floating numbers side, we just have float32 and float64

	var n1 int8 = 100   // 8 bites, goes from 00000000 -> 11111111 with a sign
	var n2 int16 = 1000 // 16 bits,
	var n3 int32 = 10000
	var n4 int64 = 100000

	// we also have int alone, using int will create a machine architectured based int

	fmt.Println(n1, n2, n3, n4)

	// these doens't support sings like -8 for example
	var un1 uint8 = 100
	var un2 uint16 = 1000
	var un3 uint32 = 10000
	var un4 uint64 = 100000

	fmt.Println(un1, un2, un3, un4)

	var floating1 float32 = 10.4
	var floating2 float64 = 110.3

	fmt.Println(floating1, floating2)

	// error is also a primitive data type created using the errors package.

	var err error = errors.New("Internal Errror")
	fmt.Println(err)
}
