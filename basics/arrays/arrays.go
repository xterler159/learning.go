package arrays

import "fmt"

func Arrays() {
	fmt.Println("====================== ARRAYS ======================")

	var t [5]int // array with 5 slots, intialized with 0 (bcz the type is int)
	t[1] = 10
	t[4] = 22
	lastIndexOfT := len(t) - 1

	var arr2 [5]float64
	var arr3 [5]string
	arr4 := [3]string{"hello", "world"}
	arr5 := [3]int{1, 2}

	fmt.Println("t:", t, ", len t:", len(t))
	fmt.Printf("t: %v\n", t)
	fmt.Println("t at index 1:", t[1])
	fmt.Println("t at index 0:", t[0])
	fmt.Println("t[lastIndexOfT]:", t[lastIndexOfT])
	fmt.Println("lastIndexOfT", lastIndexOfT)
	fmt.Println("arr2:", arr2)
	fmt.Println("arr3:", arr3)
	fmt.Println("arr3[1]:", arr3[1])
	fmt.Printf("arr4: %v\n", arr4)
	fmt.Printf("arr5: %v\n", arr5)

	fmt.Println("====================== ARRAYS ======================")
}
