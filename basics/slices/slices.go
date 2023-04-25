package slices

import (
	"fmt"
)

func Slices() {
	fmt.Println()
	fmt.Println("====================== SLICES ======================")

	s := make([]int, 3) // initializing a slice, with 3 slots
	fmt.Println("s:", s)

	s = append(s, 12)  // adding 12 into the slice
	s = append(s, 10)  // adding 10 this time into the slice
	s = append(s, 100) // adding 10 this time into the slice
	s = append(s, 99)  // adding 10 this time into the slice

	fmt.Println("after all appends:", s)
	fmt.Println("len s:", len(s))
	fmt.Println("capacity (cap) s:", cap(s))

	nums := make([]int, 2, 3)
	nums[0] = 10
	nums[1] = 44
	// nums[2] = 8743 // will not work bcz there is only 2 slots availaible, on slice init

	fmt.Println()
	fmt.Println("nums:", nums)
	fmt.Println("nums len:", len(nums))
	fmt.Println("nums cap:", cap(nums))

	nums = append(nums, 2)
	nums = append(nums, 99)

	fmt.Println()
	fmt.Println("nums after append:", nums)
	fmt.Println("len of nums after append:", len(nums))
	fmt.Println("cap of nums after append:", cap(nums))

	letters := []string{"g", "o", "l", "a", "n", "g"} // short way to init a slice

	fmt.Println()
	fmt.Println("letters:", letters)
	fmt.Println("len of letters:", len(letters))
	fmt.Println("cap of letters:", cap(letters))

	letters = append(letters, "!")

	fmt.Println()
	fmt.Println("letters after append:", letters)
	fmt.Println("len of letters after append:", len(letters))
	fmt.Println("cap of letters after append:", cap(letters))

	sub1 := letters[0:2] // sub slice, this means it will create a slice with 2 first elements
	sub2 := letters[2:]  // sub slice, this means it will create a slice starting at the index 2, and will continue until the end
	sub2[0] = "UP"

	fmt.Println()
	fmt.Println("sub1:", sub1)
	fmt.Println("sub2:", sub2)
	fmt.Println("letters:", letters) // letters will be altered because sub1 and sub2 is pointing to the letters slice

	// we can prevent this copying the array
	sub2Copy := make([]string, len(sub1))
	copy(sub2Copy, sub2)

	fmt.Println()
	fmt.Println("sub2Copy:", sub2Copy)

	fmt.Println("====================== SLICES ======================")
}
