package main

import (
	"fmt"
	"strings"
	"time"

	"learning.go/hellworld/input"
	"rsc.io/quote"
)

func main() {
	testInput := input.TestKeyboard()
	currentTime := time.Now()

	fmt.Println(strings.ToUpper(testInput))
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println("currentTime:", currentTime)
}
