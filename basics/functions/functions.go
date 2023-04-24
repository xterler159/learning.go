package functions

import "fmt"

func TestFunc(a int, b int) {
	fmt.Printf("a: %d, b: %d\n", a, b)
}

func TestFunc2(a, b float64) { // this means these 2 params will be the float64 type
	fmt.Printf("a: %f, b: %f", a, b)
	fmt.Println()
}

func SumNamedReturn(x, y, z int) (sum int) {
	sum = x + y + z
	return sum
}

func ReturnMultipleReturns() (int, string) {
	return 20, "my string"
}
