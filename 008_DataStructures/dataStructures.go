package datastructures

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func DataStructures() {
	// Array declaration
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		fmt.Printf("%d", arr[i])
	}
	fmt.Print("\n")

	arr2:= [5]int{1,2,3,4,5}
	for i := 0; i < 5; i++ {
		fmt.Printf("%d", arr2[i])
	}
	fmt.Print("\n")

	// Let interpreter define array length
	arr3:= [...]int{1,2,3,4,5}
	for i := 0; i < 5; i++ {
		fmt.Printf("%d", arr3[i])
	}
	fmt.Print("\n")

	// Matrix
	mat:= [2][2]int{{1,2}, {3,4}};
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d", mat[i][j])
		}
	}
	fmt.Print("\n")

	// ? Note: The size of the array is part of the type so an [2]int array and a [3]int array are different data types
	// ? Note 2: The = or := operators copies the right value into the left variable, so it can be expensive
	
	// Slices

	// ? Note 3: Slices are like references to memory address of a certain array
	// ? Note 4: In slices, the = and := operators makes a reference to a memory address so you can have two or more slices pointing
	// ? 		 to the same memory address
	// ? Note 5: In slices, the size of the slice is not part of the type, the type always will be []int
	// ? Note 6: The index validation will be made on run time, so the app can compile and provoke a runtime error if you try to access
	// ? 		 to a inexistent index
	// ? Note 7: Slices doesnt have fixed length, you can use append to add an element
	// ? Note 8: Slices can create sub-slices named Views
	
	// Slice declaration
	slice := []int{1,2}
	
	// var err = slice[4] // ! this is posible, but will throw a runtime error
	
	// append data
	slice = append(slice, 3,4);

	// Get slice length and capacity
	fmt.Printf("Len: %d, Cap: %d\n", len(slice), cap(slice))
	slice = append(slice, 3,4);
	fmt.Printf("Len: %d, Cap: %d\n", len(slice), cap(slice))

	// Control the initial size
	slice2:= make([]int, 0, 8)
	slice2 = append(slice2, 1);
	fmt.Printf("Len: %d, Cap: %d\n", len(slice2), cap(slice2))
	fmt.Println("")

	// Slices copy
	slice3:=make([]int, len(slice))
	copiedElements := copy(slice3, slice)

	fmt.Printf("Copied Elements: %d\nSlice: %v\n", copiedElements, slice3)

	fmt.Println("")
	
	// Slices as params
	var arrT [3]int = [3]int{1,2,3}
	var sliceT []int = []int{1,2,3}

	fmt.Printf("Array: %v\nSlice: %v\n", arrT, sliceT)

	func (vec [3]int, sli[]int)  {
		vec[0] = 0;
		if len(sli) > 0 {
			sli[0] = 0;
		}
	}(arrT, sliceT)
	fmt.Println("")

	fmt.Printf("Array: %v\nSlice: %v\n", arrT, sliceT)

	// Loop through an array/slice
	for i := 0; i < len(sliceT); i++ {
		fmt.Println(i)
	}
	fmt.Println("")

	for i, value := range sliceT {
		fmt.Println("[", i, "]", value)
	}
	fmt.Println("")

	for _, value := range sliceT {
		fmt.Println(value)
	}
	fmt.Println("")

	// Creating views from slices
	// syntax slice[<begining> : <ending>]
	// Slices counts with a length of 6, and we are slicing from the index n to n-1
	view := slice[0:2] // From 0 to 1 

	for _,val := range view{
		fmt.Println(val)
	}
	fmt.Println("")

	view = slice[1:4] // From 1 to 3

	for _,val := range view{
		fmt.Println(val)
	}
	fmt.Println("")

	// ? On a view declaration the begining value will be 0 and the end value will be the length of the slice
	// ? A view is also a reference of its parent, so we can modify a slice from a view

	// Variable args usage
	fmt.Println(add(1,2,3,4))
	fmt.Println(add(slice...))

	//mix slices
	fmt.Println("View: ", view, "Slice: ",slice2)
	view = append(view, slice2...)

	fmt.Println(view, view)

	// string length
	str := "Hello, world"

	fmt.Println("Str Len: ", utf8.RuneCountInString(str))

	//str to slice
	strSlice := []rune(str)
	strView := strSlice[:5]

	fmt.Printf("OriginalString: %s\nRune Slice: %v\nRune View: %v\n", str, strSlice, strView)
	// Modifying string 
	strView[0] = 'h';
	fmt.Printf("Rune Slice: %v\nString Modified: %s\n", strSlice, string(strSlice))

	// The correct way for creating strings
	var sb strings.Builder;

	sb.WriteString("Hello\n")
	sb.WriteString("This is the best way to create strings\n")
	sb.WriteString("But you can only create, you cant erease\n")
	sb.WriteString("You can convert numbers to string with strconv.Itoa")
	sb.WriteByte('(')
	sb.WriteString(strconv.Itoa(10))
	sb.WriteByte(')')

	fmt.Println(sb.String())
}

// variable args functions

func add(nums ...int) int{
	var total int =0

	for _, num := range nums{
		total += num
	}

	return total
}