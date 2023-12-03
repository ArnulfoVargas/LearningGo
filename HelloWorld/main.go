package main

import "fmt"

func main() {
	// Comment
	/* Also a comment */
	fmt.Println("Hello world")
	
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
	
	fmt.Println(hello)
	fmt.Println(text)
}