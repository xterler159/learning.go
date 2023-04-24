package switchstatement

import "fmt"

func Switch(testVar int) {
	fmt.Println()
	fmt.Println("====================== SWITCH STATEMENT ======================")

	switch testVar {
	case 1:
		fmt.Println("1")

	case 2:
		fmt.Println("2")

	//case testVar > 2: this case will not work
	//fmt.Println("")
	default:
		fmt.Println("Default statement 1")
	}

	switch {
	case testVar%2 == 0: // this will work
		fmt.Println("dividable by 0")

	default:
		fmt.Println("Default statement 2")
	}

	fmt.Println("Switch func")
	fmt.Println("====================== SWITCH STATEMENT ======================")
	fmt.Println()
}
