package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// maps are key-value pairs structure where the keys and the values are strictly typed.
	u1 := map[string]string{
		"name": "Richard",
		"age":  "32",
	}
	fmt.Println(u1)

	// we can also have nested maps

	u2 := map[int8]map[string]string{
		1: {"name": "Richard", "age": "32"},
		2: {"name": "Charles", "age": "28"},
	}

	fmt.Println(u2)

	// deleting a entry in a map
	delete(u2, 1)
	fmt.Println(u2)

	// adding a entry into a map
	u2[1] = map[string]string{"name": "Richard", "age": "32"}
	fmt.Println(u2)

	// edit a entry in a map
	u2[1]["age"] = "33"
	fmt.Println(u2)
}
