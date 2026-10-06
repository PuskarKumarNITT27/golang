package main

import (
	"fmt"
	"strings"
)

func main(){
	firstName:= "puskar "
	lastName:= "Kumar"

	fullName := firstName + " " + lastName

	fmt.Println(fullName)

	fmt.Println(strings.ToUpper(fullName))
}