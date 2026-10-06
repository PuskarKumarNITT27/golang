package main

import "fmt"

func main(){
	views1 := 1000.0
	views2 := 2000.0

	totalViews := views1 + views2
	

	likes := 10
	likes++
	likes++

	avgViews := totalViews / 7

	fmt.Println("TotalViews, avgViews ,likes", totalViews,avgViews,likes)

}