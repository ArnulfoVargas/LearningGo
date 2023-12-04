package main

import (
	"fmt"
)

import (
	"variables"
	"inputs"
	"conditionals"
	"loops"
	"pointers"
	"functions"
	)

func main() {
	// Comment
	/* Also a comment */
	fmt.Println("Hello world")

	variables.VariablesTest()

	inputs.InputsTest()

	conditionals.ConditionalsTest()

	loops.LoopsTest()

	pointers.PointersTest()

	functions.FunctionsTest()
}