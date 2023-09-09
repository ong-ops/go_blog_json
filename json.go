package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	message := Message{"Alice", "Hello", 1294706395881547000}

	fmt.Println("\n=============\nJSON Encoding\n=============")
	jsonMessage, err := json.Marshal(message)
	fmt.Println(jsonMessage, err)

	fmt.Println("\n=============\nJSON Decoding\n=============")
	decodedErr := json.Unmarshal(jsonMessage, &message)
	fmt.Println(message, decodedErr)

	fmt.Println("\n=============\nJSON Decoding - Unexactly match\n=============")
	unExactMatchJson := []byte(`{"Name":"Bob","Food":"Pickle"}`)
	var unExactMatchMessage Message
	unExactMatchErr := json.Unmarshal(unExactMatchJson, &unExactMatchMessage)
	fmt.Println(unExactMatchMessage, unExactMatchErr)

	fmt.Println("\n=============\nGeneric JSON with interface\n=============")
	var i interface{}
	i = "a string"
	i = 2011
	i = 2.777
	r := i.(float64)
	fmt.Println("the circle's area", math.Pi*r*r)

	switch v := i.(type) {
	case int:
		fmt.Println("twice i is", v*2)
	case float64:
		fmt.Println("the reciprocal of i is", 1/v)
	case string:
		h := len(v) / 2
		fmt.Println("i swapped by halves is", v[h:]+v[:h])
	default:
		// i isn't one of the types above
	}

	fmt.Println("\n=============\nJSON Decoding - Arbitrary data\n=============")
	arbitraryJson := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var arbitraryInterface interface{}
	arbitraryErr := json.Unmarshal(arbitraryJson, &arbitraryInterface)
	fmt.Println(arbitraryInterface, arbitraryErr)
	arbitraryMessage := arbitraryInterface.(map[string]interface{})
	fmt.Println(arbitraryMessage)

	for key, value := range arbitraryMessage {
		switch vv := value.(type) {
		case string:
			fmt.Println(key, "is string", vv)
		case float64:
			fmt.Println(key, "is float64", vv)
		case []interface{}:
			fmt.Println(key, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(key, "is of a type I don't know how to handle")
		}
	}

	fmt.Println("\n=============\nJSON Decoding - Reference Types\n=============")
	type FamilyMember struct {
		Name    string
		Age     int
		Parents []string
	}
	var member FamilyMember
	memberErr := json.Unmarshal(arbitraryJson, &member)
	fmt.Println(member, memberErr)
}
