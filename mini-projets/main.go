package main

import (
	"fmt"

	goplace "learning.go/mini-projets/Goplace"
)

func main() {
	_, err := goplace.FindReplaceFile("./tests.txt", "hello", "ok")

	if err != nil {
		fmt.Println("Error while calling FindReplaceFile function:", err)
	}

	// fmt.Println("occ:", occ)

	// count := strings.Count("hello hello hello", "hello")
	// fmt.Println(count)
}
