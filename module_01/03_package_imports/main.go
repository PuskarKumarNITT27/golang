package main

/*
 => import brings external packages into the file that you are working on 
 => where actually you need it 
*/


import (
	"fmt"
	"math"
)

func main(){

	// packageName.function -> calls a function from the required package
	
	fmt.Println("hell0")
	fmt.Println(math.Sqrt(34))
}