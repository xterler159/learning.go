package rangestatement

import "fmt"

func RangeStatement() {
	fmt.Println()
	fmt.Println("====================== RangeStatement, RangeStatement func======================")

	// basic functioning range
	for index, value := range []int{0, 10} {
		fmt.Println("index:", index, "value:", value)
	}
	fmt.Println()

	// using range to on a strings slice
	names := []string{"Bob", "Alice", "Kevin", "Michel"}

	for index, name := range names {
		fmt.Printf("Username=%s (index=%d)\n", name, index)
	}
	fmt.Println()

	// looping through a string
	for _, char := range "golang" {
		// this will print the value in bytes
		fmt.Printf("char: %v\n", char) // %v is for value

	}

	fmt.Println()
	for _, char := range "golang" {
		// this will print the value converted to a string
		fmt.Printf("char (converted) %v\n", string(char))
	}

	fmt.Println("====================== RangeStatement, RangeStatement func======================")
}
