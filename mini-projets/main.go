package main

import (
	"fmt"

	goplace "learning.go/mini-projets/Goplace"
)

func main() {
	occ, lines, err := goplace.FindReplaceFile("./tests.txt", "hello", "ok")

	if err != nil {
		fmt.Println("Error while calling FindReplaceFile function:", err)
		return
	}

	fmt.Println("occ:", occ)
	fmt.Println("lines:", lines)
}
