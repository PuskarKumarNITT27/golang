package main

import (
	"fmt"
)

func main() {
	isLogged := true // inferred as boolean type
	isAdmin := true
	hasSubscription := false

	// AND &&
	canOpenDashboard := isLogged && hasSubscription

	canDeletePost := isAdmin || (isLogged && hasSubscription)

	fmt.Println(canOpenDashboard, canDeletePost)

	age := 2
	isAdult := age > 18
	fmt.Println(isAdult)
}
