package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}
func main() {
	// Separating this function because it's a side effect.
	// Printing Hello World is our domain. fmt.Println is outside domain.
	fmt.Println(Hello("Daniel"))
}
