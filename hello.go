package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const english_greeting = "Hello, "
const spanish_greeting = "Hola, "
const french_greeting = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return checkLanguage(language) + name
}

//something to remember. functions that start with Upper case are public. Lower case means private.

func checkLanguage(language string) (greeting string) {

	switch language {
	case spanish:
		greeting = spanish_greeting
	case french:
		greeting = french_greeting
	default:
		greeting = english_greeting
	}

	return
}

func main() {
	// Separating this function because it's a side effect.
	// Printing Hello World is our domain. fmt.Println is outside domain.
	fmt.Println(Hello("Daniel", ""))
}
