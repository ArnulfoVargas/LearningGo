package pointers

import "fmt"

func PointersTest() {
	// Pointers only point to variables with the same data type
	var num uint8 = 10
	var pointer1 *uint8 // nil pointer (null pointer)

	if pointer1 == nil {
		fmt.Println("Nil")
	}

	pointer1 = &num // pointing to num memory address

	fmt.Println(*pointer1)
	*pointer1 = 20;
	fmt.Println(num);

	var num2 uint8= 20;
	var pointer2 *uint8 = &num2;

	if pointer1 == pointer2 {
		fmt.Println("Pointing to the same address")
	} else {
		fmt.Println("Different address")
	}

	if *pointer1 == *pointer2 {
		fmt.Println("Same value")
	} else {
		fmt.Println("Different Value")
	}
}
