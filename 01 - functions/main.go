package main

import (
	"fmt"
	"log"
)

type shit struct {
	shittyProperty string
}

func main() {
	shittyVoid()

	shittyWithMultipleReturnTypes()

	shittyWithOneReturnValueOfString()

	shittyValue := shit{shittyProperty: "Shit"}

	shittyWithAnInput(shittyValue)

	// function literal or anonymuse function
	(func() {
		log.Print("I'm shitty function literal")
	})()

	anotherShittyfunctionLiteral := (func() {
		log.Print("I'm another shitty function literal")
	})

	anotherShittyfunctionLiteral()

}

// define a void function
func shittyVoid() {
	log.Println("Hey I'm shitty void function without any return value")

	shittyString := shittyWithOneReturnValueOfString()

	fmt.Println(shittyString)
}

func shittyWithOneReturnValueOfString() string {
	log.Println("Hey I'm shitty function with a return value of type string")

	return "Shitty value"
}

func shittyWithMultipleReturnTypes() (string, int, []byte) {
	log.Println("Hey I'm shitty function multiple return type")
	shittyString := "shitty value"
	shittyByteSlice := []byte{uint8(1)}
	shittyInt := 8

	return shittyString, shittyInt, shittyByteSlice
}

func shittyWithAnInput(shittyInput shit) bool {
	log.Println("Hey I'm shitty function with an input param")
	log.Println(shittyInput)
	return true
}
