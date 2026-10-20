package main

import "fmt"

func sumAll(nums ...int) int {
	total := 0

	for _, currentValue := range nums {
		total = total + currentValue
	}

	return total
}

func main() {

	fmt.Println(sumAll(1, 2, 3, 4, 5))

	values := []int{10, 23}
	fmt.Println(sumAll(values...))

	res := func(n int) int {
		return n * 2
	}

	fmt.Println(res(2))

	// IIFE

	res1 := func(a int, b int) int {
		return a + b
	}(5, 10)

	fmt.Println(res1)

}
