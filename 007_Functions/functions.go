package functions

import (
	"fmt"
	"math/rand"
)

// Syntax:
// func <name>() { body }

func FunctionsTest() {
	hello()
	// Calling functions with parameters
	hello2("Arny", "Vargas")
	hello3("Arny", "Vargas")
	fmt.Println(add(4, 5))

	// Multiple return usage
	var x, y int = maxMin(5, 7)
	fmt.Println(x, y)

	// flip variable values
	x, y = y, x

	fmt.Println(x, y)

	// usage doesnt change
	x, y = maxMin2(5, 7)

	fmt.Println(x, y)

	// If you wont use some value you can use an empty identifier
	max, _ := maxMin2(10, 25)
	fmt.Println(max)

	// usage reference func
	addOne(&max)
	fmt.Println(max)

	// function literals
	var randomNum func() int
	randomNum = func() int {
		return rand.Int()
	}

	fmt.Println(randomNum())

	// Using a function for a function literal assigment
	operation := add
	fmt.Println(operation(3, 5))
	operation = multiply
	fmt.Println(operation(3, 5))

	// Invocation on declaration
	func() {
		fmt.Println("Invoked")
	}()
}

// the function will be private if you begin the naming with lowercase
func hello() {
	fmt.Println("Hello 1, user")
}

// Parameters
// Syntax:
// func <name>(param1 type...) { body }
func hello2(name string, lastName string) {
	fmt.Println("Hello 2,", name, lastName)
}

// If multiple parameters share the same data type, you can avoid writting it except for the last one
func hello3(name, lastName string) {
	fmt.Println("Hello 3,", name, lastName)
}

// Return
// Syntax:
// func <name>(param type...) return type { body }
func add(num1, num2 int) int {
	return num1 + num2
}

func multiply(num1, num2 int) int {
	return num1 * num2
}

// multiple returns
// Syntax:
// func <name>(param type...) (return type1, return type 2...) { body }
func maxMin(num1, num2 int) (int, int) {
	if num1 > num2 {
		return num1, num2
	}

	return num2, num1
}

// Multiple named values
// Syntax:
// func <name>(param type...) (<name> type1, <name> type 2...) { body }
func maxMin2(num1, num2 int) (max int, min int) {
	if num1 > num2 {
		max = num1
		min = num2
	} else {
		max = num2
		min = num1
	}
	return
}

// Reference parameter
func addOne(n *int) {
	*n += 1
}

// !!! GO DOESNT SUPPORT METHOD OVERLOAD
