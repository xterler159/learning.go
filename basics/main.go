package main

import (
	"fmt"
	"time"

	"learning.go/basics/ifelse"
)

func main() {
	// variables
	var age int = 20
	fmt.Println(age)

	var weight, height = 90, 190
	fmt.Println(weight, height)

	var t = time.Now()
	fmt.Println(t)

	email := "test@test.com"
	fmt.Println(email)

	// control structure
	ifelse.IfElse(15)
	fmt.Printf("user age: %d\n", 2)
}
