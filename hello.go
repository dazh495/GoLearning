package main

import "fmt"

const greeting = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return greeting + name
}
func main() {
	// Separating this function because it's a side effect.
	// Printing Hello World is our domain. fmt.Println is outside domain.
	fmt.Println(Hello("Daniel"))
}
