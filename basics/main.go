package main

import (
	"fmt"
	"time"

	"learning.go/basics/arrays"
	"learning.go/basics/functions"
	"learning.go/basics/ifelse"
	"learning.go/basics/loops"
	"learning.go/basics/rangestatement"
	"learning.go/basics/slices"
	"learning.go/basics/switchstatement"
	"learning.go/basics/typeconversion"
)

func main() {
	// variables
	var age int = 20
	fmt.Println(age)

	var weight, height = 90, 190
	fmt.Println(weight, height)

	var t = time.Now()
	fmt.Println(t)

	email := "test@test.com"
	fmt.Println(email)

	// control structure
	ifelse.IfElse(15)
	switchstatement.Switch(3)

	// types conversion - basics
	typeconversion.TypeConversion()

	// functions
	fmt.Println("====================== FUNCTIONS, TestFunc======================")
	functions.TestFunc(20, 80)
	functions.TestFunc2(12, 89.0)
	retValue := functions.SumNamedReturn(10, 20, 30)
	fmt.Printf("retValue %d\n", retValue)

	val1, val2 := functions.ReturnMultipleReturns()
	_, val2 = functions.ReturnMultipleReturns() // ignoring the first returned value
	// testVar, val2 = functions.ReturnMultipleReturns() // this will not work bcz testVar is not declared (=, not :=)

	fmt.Printf("val1: %d, val2: %s\n", val1, val2)
	fmt.Printf("val2: %s\n", val2)

	fmt.Println("====================== FUNCTIONS, TestFunc======================")

	// arrays
	arrays.Arrays()

	// slices
	slices.Slices()

	// loops
	loops.Loops()

	// range
	rangestatement.RangeStatement()
}
