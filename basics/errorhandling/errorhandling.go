package errorhandling

import (
	"fmt"
	"io/ioutil"
)

func ErrorHandling() {
	fmt.Println()
	fmt.Println("====================== ErrorHandling ======================")

	_, err := errorDemo()

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// example on reading a file
	data, err := readFile("./errorhandling/test.txt")

	if err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
		fmt.Println("====================== ErrorHandling ======================")

		return
	}
	fmt.Printf("File content: %v\n", data)

	fmt.Println("====================== ErrorHandling ======================")
}

func errorDemo() (int, error) {
	return 1, nil
}

func readFile(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	if len(data) == 0 {
		// return "", errors.New("Empty content in file")
		return "", fmt.Errorf("Empty content in (filename=%v)", fileName) // formatting error message
	}

	return string(data), nil
}
