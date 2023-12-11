package structs

import (
	"fmt"
)

//Struct declaration
type Cuboid struct{
	width float64
	height float64
	deep float64
}

// Functionality
func (c Cuboid) Volume() float64 {
	return c.deep * c.height * c.width
}

func (c *Cuboid) SetToCero(){
	c.deep = 0
	c.height = 0
	c.width = 0
}

func (c Cuboid) String() string {
	return fmt.Sprintf("Cuboid w: %v x h: %v x d: %x", c.width, c.height, c.deep)
}

func StructsTest(){
	// Cuboid instance with default values
	c1 := Cuboid{}
	var c2 Cuboid;

	// properties modification
	c1.deep = 10
	c2.height = 20

	// Cuboid instance with custom values
	c3 := Cuboid{
		width: 10,
		height: 20,
		deep: 30,
	}

	fmt.Printf("%f\n", c3.Volume())
	c3.SetToCero()
	fmt.Printf("%#v\n",c3) 
	fmt.Printf("%s\n",c3.String())

	structAppend()
	constructorTest()
}

// Struct append
// Go doesnt allow inheritance because go is not a OOP language
// But we can append a struct into another struct
type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) BaseArea() float64{
	return r.height * r.width
} 

type RectangularPrism struct{
	Rectangle
	deep float64
}

func structAppend(){
	// Using rectangle data and funcs into prism
	prism := RectangularPrism{}
	prism.width = 10;
	prism.height = 20;
	prism.deep = 30

	fmt.Println(prism.BaseArea())

	// A complete declaration
	prism2 := RectangularPrism{
		Rectangle: Rectangle{
			width: 10,
			height: 20,
		},
		deep: 30,
	}
	fmt.Printf("%#v\n", prism2)
}

// Creating a Custom Constructor
type Option func(*RectangularPrism)

func width(width float64) Option{
	return func(rp *RectangularPrism) {
		rp.width = width;
	}
}

func height(height float64) Option{
	return func(rp *RectangularPrism) {
		rp.height = height;
	}
}

func deep(deep float64) Option{
	return func(rp *RectangularPrism) {
		rp.deep = deep;
	}
}

func prismConstructor(options ...Option) RectangularPrism {
	prism := RectangularPrism{}

	for _, option := range options{
		option(&prism)
	}

	return prism
}

// Using Constructor
func constructorTest(){
	prism := prismConstructor(width(15), deep(25))

	fmt.Printf("%#v", prism)
}