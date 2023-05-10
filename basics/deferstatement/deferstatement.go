package deferstatement

import (
	"fmt"
)

func start() {
	fmt.Println("Start func")
}

func finish() {
	fmt.Println("Finish func")
}

func DeferStatement() {
	fmt.Println()
	fmt.Println("====================== DeferStatement, DeferStatement func ======================")

	defer start()
	finish()

	names := []string{"Jean", "Kevin", "Michelle", "Louise"}
	for _, name := range names {
		fmt.Println("Hello", name)
		defer fmt.Println("Goodbye", name) // defer is working on a FIFO (first in, first out)
	}

	fmt.Println("====================== DeferStatement, DeferStatement func ======================")
}
