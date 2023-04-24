package typeconversion

import (
	"fmt"
)

func TypeConversion() {
	fmt.Println("====================== TYPE CONVERSION ======================")

	ten := "10"
	fmt.Println("ten + '2'", ten+"2")

	// var x int = ten + 1 // can't because go is strongly typed (typage fort)

	intValue := 20
	// convertedValue := intValue + 0.2 // will cause a compilation error (strongly typed)
	convertedValue := intValue + 10 // will work
	// convertedValue = float32(intValue) + 0.5 // will not work
	n := float32(intValue) + 0.5 // will work. We have to create a new variable in order to do the conversion

	var percentage float64 = 2.9
	percentage = percentage + 20

	fmt.Println("convertedValue:", convertedValue)
	fmt.Println("n:", n)
	fmt.Println("percentage:", percentage)
	fmt.Println("percentage int(percentage):", int(percentage)) // we will have a data loss

	a := 42
	b := float64(a) + 0.42

	fmt.Println("b:", b)
	fmt.Println("====================== TYPE CONVERSION ======================")
	fmt.Println()
}
