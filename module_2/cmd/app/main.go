package main

import (
	"fmt"
	"go-module/internal/greet"
)


func main(){
	
	name := greet.Hello("Puskar")

	fmt.Println(name)
}