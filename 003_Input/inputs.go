package inputs

import "fmt"

func InputsTest(){
	// Input
	var age int
	fmt.Printf("Age?")
	fmt.Scanln(&age)
	fmt.Printf("Your age is: %d\n", age)

	var hours, minutes, seconds int
	fmt.Printf("HH:MM:SS?")
	fmt.Scanf("%d:%d:%d\n", &hours, &minutes, &seconds)
	fmt.Printf("%d:%d:%d\n", hours, minutes, seconds)
}