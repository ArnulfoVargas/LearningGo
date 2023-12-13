package interfaces

import "fmt"

// Interfaces in go are easier
// You only need to declare an interfaces
type Greet interface{
	// Now we need to declare the interface functions
	GreetPeople()
	GoodBye()
}

// We dont need some kind of inheritance, just having the functions of the interface
// will be part of the interface
type Human struct{}

func (h Human) GreetPeople() {
	fmt.Println("Hello")
}

func (h Human) GoodBye() {
	fmt.Println("Bye bye")
}
// Interfaces in Go has Duck typing
// "If it walks like a duck and it quacks like a duck, then it must be a duck"

func InterfacesTest(){
	human:=Human{}

	human.GreetPeople()
	human.GoodBye()

	var human2 Greet = Human{}  // Using the interface 
	human2.GoodBye()
}