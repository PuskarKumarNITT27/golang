package main

import "fmt"

// struct groups related fields into one type

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {

	u1 := User{
		ID:    1,
		Name:  "Sangam",
		Email: "sangam@gmail.com",
		Age:   100,
	}

	fmt.Println(u1, u1.ID, u1.Email)

	// mutable by default
	u1.Age = 200

	fmt.Println(u1)

	u2 := User{
		Name:  "john",
		Email: "dsfsdfs",
	}

	fmt.Println("partial user", u2)

}
