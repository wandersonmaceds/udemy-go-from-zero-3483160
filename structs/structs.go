package main

import "fmt"

type user struct {
	name string
	age  int8
}

type address struct {
	street string
	number int8
}

type userWithAddress struct {
	name    string
	age     int8
	address address
}

type student struct {
	user
	course string
}

func main() {
	fmt.Println("Dealing with complex data")

	// we can create struct and assign each value later
	var u user
	u.name = "Richard"
	u.age = 32

	fmt.Println(u)

	// we can also create and assign the values directly by using {}
	u2 := user{"Richard", 32}
	fmt.Println(u2)

	// in case we need to fill just one field, we can use named parameters
	// non-initialized fields are set to the zero value of each type
	// strings are empty and numbers are 0.

	u3 := user{name: "Richard"}
	fmt.Println(u3)
	fmt.Println(user{}) //empty struct

	// we can also have nested structs
	u4 := userWithAddress{
		name: "Richard",
		age:  32,
		address: address{
			street: "John Doe Street",
			number: 6,
		},
	}

	fmt.Println(u4)

	newAddress := address{"John Doe Street", 6}
	u5 := userWithAddress{"Richard", 32, newAddress}
	fmt.Println(u5)

	// student acts like a inheritance, copying all keys from user, without the need to use it a nested struct
	fmt.Println("------------------")
	fmt.Println(student{})
	e1 := student{user{"Richard", 32}, "Engineering"}
	fmt.Println(e1, e1.name, e1.course)

}
