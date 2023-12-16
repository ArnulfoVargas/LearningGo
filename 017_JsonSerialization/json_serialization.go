package jsonSerialization

import (
	"encoding/json"
	"fmt"
)

type data struct {
	Name string
	Age  uint
	Mail string
}

func SerializationTest() {
	// Struct to Json
	user := data{
		Name: "NameTest",
		Age:  18,
		Mail: "test@test.com",
	}
	jsonResult, err := json.Marshal(user);

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonResult))
	}

	user2 := data{}

	if err := json.Unmarshal(jsonResult, &user2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", user2)
	}

	// Storing in map
	user3 := make(map[string]interface{})

	err = json.Unmarshal(jsonResult, &user3)

	if err != nil{
		fmt.Println(err)
	} else {
		for k,v := range user3{
			fmt.Printf("[%s] %v\n", k, v)
		}
	}

	// Storing in slice/array
	arrJson := []byte(`["Hello", " ", "World"]`)
	var parsedArr []string

	err = json.Unmarshal(arrJson, &parsedArr)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Printf("%v", parsedArr)
	}
}