package loops

import "fmt"

func Loops() {
	fmt.Println()
	fmt.Println("====================== LOOPS ======================")

	sum := 0

	for i := 0; i < 10; i++ { // classique loop
		fmt.Println("loop:", i)
		sum += 1
	}

	fmt.Println("sum:", sum)

	eventCount := 0

	for eventCount < 4 {
		fmt.Println("Retrieving events ...")
		eventCount++

		if eventCount == 3 {
			fmt.Printf("Got %d notifications, updating phone notif\n", eventCount)
		}
	}

	fmt.Println()

	// "while" loop
	i := 0
	for {
		i++

		if i%2 != 0 {
			fmt.Println("Odd number")
			continue
		}

		fmt.Println("Looping...")

		if i >= 10 {
			fmt.Println("Loop end")
			break
		}
	}

	fmt.Println("====================== LOOPS ======================")
}
