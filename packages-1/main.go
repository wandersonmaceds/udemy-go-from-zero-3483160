package main

import (
	"fmt"
	"module/utils"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writting from main!")
	utils.Write()
	error := checkmail.ValidateFormat("123")
	noerror := checkmail.ValidateFormat("user@mail.com")

	fmt.Println(error)
	fmt.Println(noerror)
}
