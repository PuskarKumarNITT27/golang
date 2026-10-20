package main

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "Sangam"
	lastName := "Mukherjee"
	fullName := firstName + " " + lastName

	fmt.Println(fullName)

	fmt.Println(strings.ToUpper(fullName))
}
