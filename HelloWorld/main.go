package main

import (
	"fmt"
	"math/rand"
)

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

	// Input
	var age int
	fmt.Printf("Age?")
	fmt.Scanln(&age)
	fmt.Printf("Your age is: %d\n", age)

	var hours, minutes, seconds int
	fmt.Printf("HH:MM:SS?")
	fmt.Scanf("%d:%d:%d\n", &hours, &minutes, &seconds)
	fmt.Printf("%d:%d:%d\n", hours, minutes, seconds)

	//Conditionals
	randomNum := rand.Int()
	fmt.Printf("%d\n", randomNum)
	if randomNum%2 == 0 {
		fmt.Println("Even")
	} else { // needs to be next to the if end
		// In this line will give you an error
		fmt.Println("Odd")
	}
	// Variable declaraction on if
	if randomNum2 := rand.Int(); randomNum2%2 == 0 {
		fmt.Println("Even", randomNum2)
	} else {
		fmt.Println("Odd", randomNum2)
	}
	// randomNum2 no longer exists at this point

	// Switch
	var letter rune = 'a'
	switch letter {
	case 'a':
		fmt.Printf("vowel")
	case 'A':
		fmt.Printf(" uppercase\n")
	}
	// Default case
	switch letter2 := 'B'; letter2 {
	case 'b':
		fmt.Printf("vowel")
	case 'B':
		fmt.Printf(" uppercase\n")
	default:
		fmt.Printf("Unrecognized\n")
	}

	// Multiple comprobations
	switch letter2 := 'B'; letter2 {
	case 'A', 'a', 'E', 'e', 'I', 'i', 'O', 'o', 'U', 'u':
		fmt.Printf("vowel\n")
	default:
		fmt.Printf("Other\n")
	}

	// Complex comprobations
	num3 := 25
	switch {
	case num3 > 0 && num3 < 10:
		fmt.Printf("Less than 10\n")
	case num3 > 10 && num3 < 20:
		fmt.Printf("Less than 20\n")
	case num3 > 20 && num3 < 30:
		fmt.Printf("Less than 30\n")
	default:
		fmt.Printf("I dont know\n")
	}

	// switch cascade
	num4 := 5
	switch {
	case num4 > 0 && num4 < 10:
		fmt.Printf("Less than 10\n")
		fallthrough
	case num4 > 10 && num4 < 20:
		fmt.Printf("Less than 20\n")
		fallthrough
	case num4 > 20 && num4 < 30:
		fmt.Printf("Less than 30\n")
	default:
		fmt.Printf("I dont know\n")
	}

	// Loops
	
	// Kind of loop in rust or While true in other languajes
	for {
		fmt.Print("Close? (y/n): ")
		var c rune;
		
		fmt.Scanf("%c\n", &c)

		if c == 'n' || c == 'N'{
			continue
		}
		if c == 'y' || c == 'Y'{
			break
		}
		fmt.Println("Not recognized")
	}

	fmt.Println("First loop closed")

	var c rune;
	for c != 'y' && c != 'Y'{
		fmt.Print("Close? (y/n): ")
		fmt.Scanf("%c\n", &c)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}