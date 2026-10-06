package main

import "fmt"

func main(){
	var city string 
	city = "London"

	var channel = "puskar" // inferred to string 

	// shorthand nottaion 
	subs := 10000
	subs += 100

	fmt.Println(city,channel,subs)


	likes, comments := 10000,20
	fmt.Println(likes,comments)
}