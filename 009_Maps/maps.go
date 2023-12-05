package mapsObject

import (
	"fmt"
	"math/rand"
)

func MapsTest() {
	//Full declaration
	map1 := map[int]string{
		1: "Hello",
		2: "World",
		3: "Bonjour",
		4: "Aloha",
	}
	// Simple declaration
	map2 := map[int]string{}

	// Overwrite
	// Syntax map[key] = value
	map1[1] = "Goodbye"

	// Add
	map2[2] = "Chicago"

	// Read
	// Syntax map[key] -> return value
	fmt.Println(map2[2])

	//validate
	// Syntax map[key] -> return value, containsKey
	word, ok := map2[1];

	if ok {
		println(word)
	}else{
		println("Word not found")
	}

	//delete entries
	// Syntax delete(<map>, <key>)
	delete(map1, 2)

	// Loop through a map
	for key, value := range map1{
		fmt.Println("Key:", key, "\t", "Value:", value)
	}

	// only keys
	for key := range map1{
		fmt.Println("Key:", key)
	}

	// only values
	for _, value := range map1{
		fmt.Println("Value:", value)
	}

	// map length
	fmt.Printf("%d\n\n", len(map1))

	// Sets
	// Go doesnt support sets, but we can approach the result using maps
	set := map[int]struct{}{};

	for len(set) < 6 {
		n := rand.Intn(49)
		if _, ok := set[n]; !ok {
			set[n] = struct{}{}
			fmt.Println(n)
		}
	}
}