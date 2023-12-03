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
	
	fmt.Println(hello)
	fmt.Println(text)
}