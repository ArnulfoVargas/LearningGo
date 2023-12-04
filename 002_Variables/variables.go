package variables

import "fmt"

func VariablesTest() {
	// Long variable declaration
	var hello string = `Hello
Bonjour
Hola`

	// Short variable declaration
	text := `Hello
Bonjour
Hola`

	// The := operator can be used only for initial value declarations
	// If you declare a variable without a initial value, the value will be the default value
	// int/float = 0
	// string = ""
	// bool = false

	text = "No greetings"

	// Data casting
	var num uint8 = 10
	var num2 int = int(num)

	fmt.Println(num2)
	fmt.Println(hello)
	fmt.Println(text)

	// Constants
	const PI float32 = 3.14159
	// Multiple constants
	const (
		e   float32 = 2.71828
		g   float32 = 9.81
		tau float32 = 6.28318530717958647692
	)

	// Num basis
	// binary
	var bin int = 0b_0010_1010
	var oct int = 02754
	var hex int = 0xaf

	fmt.Printf("bin: %v\noct: %v\nhex: %v\n", bin, oct, hex)
}