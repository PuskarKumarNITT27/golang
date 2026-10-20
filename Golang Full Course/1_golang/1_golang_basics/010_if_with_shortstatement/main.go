package main

import "fmt"

func main() {

	items := 3
	pricePerItem := 4

	if total := items * pricePerItem; total >= 100 {
		fmt.Println("Eligible for shipping")
	} else {
		fmt.Println("Not Eligible for shipping ***")
	}
}
