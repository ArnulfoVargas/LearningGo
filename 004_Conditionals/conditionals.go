package conditionals

import 	(
		"fmt"
		"math/rand"	
		)

func ConditionalsTest(){
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
} 