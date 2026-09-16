package utils

import "fmt"

// the Write functions only writes a dummy messate on the default output
// uppercased functions are exported by default to be used in other packages of the module
func Write() {
	fmt.Println("Writting from Utils 1")
	write()
}
