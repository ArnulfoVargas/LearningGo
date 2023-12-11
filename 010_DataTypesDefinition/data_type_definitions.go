package datatypesdefinition

import (
	"fmt"
)

func DataTypesDefinitionTest() {
	// making a private type named age with the same properties as uint8
	type age uint8
	var ageNum uint8 = 10
	// Despite of age being a renaming of uint8, we still need to make a data cast
	var ageValue age = age(ageNum)

	fmt.Println("Age: ", ageValue)

	// Data types from slices
	type year int;
	type births []year;

	count := births{1979, 1983, 1965}
	count = append(count, 1990)
	count = append(count, 1955)

	sum := year(0)

	for _, y := range count{
		sum += y;
	}

	average := sum / year(len(count))
	fmt.Println("The average is: ",average)

	// Data types from Maps
	type NameBirth map[string]year;

	artist := NameBirth{
		"Van Gogh"	:	1853,
		"Elvis"		:	1935,
		"Dali"		: 	1904,
	}
	artist["Rick"] = 1966;

	for name, year := range artist{
		fmt.Printf("Name:%s\t\t\tYear:\t%d\n", name, year)
	}
	
	// testing Generator
	counter := Count()
	counter2 := Count()

	println("Counter 1: ", counter())
	println("Counter 1: ", counter())
	println("Counter 2: ", counter2())
	println("Counter 2: ", counter2())
	println("Counter 1: ", counter())

	// Testing Data type expansion
	ph1 := PH(3)
	fmt.Println(ph1.Category())
	ph2 := PH(8)
	fmt.Println(ph2.Category())
	ph3 := PH(7)
	fmt.Println(ph3.Category())

	count1 := Counter(0)
	count1.AddOne()
	println(count1)
	count2 := Counter(2)
	count2.AddOne()
	println(count2)
	count2.Reset(0)
	println(count2)
}
// Data types from funcs and functionale programing
type Generator func() int;

func Count() Generator {
	count := 0

	return func() int {
		count++
		return count
	}
}

// Expanding Data types
type PH float32;
func (p PH) Category() string {
	switch {
	case p < 7:
		return "Acid"
	case p > 7:
		return "Basic"
	default:
		return "Neutral"
	}
}

// Expancions that modify the value must be pointers
type Counter uint;

func (c *Counter) AddOne(){
	*c++;
}

func (c *Counter) Reset(newValue uint){
	*c = Counter(newValue);
}