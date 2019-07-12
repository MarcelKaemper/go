package main

import "fmt"

func main() {
	var x,y int
	var a,b = "Hello", "World"
	const pi = 3.14

	x = 5
	y = 10

	// Short variable declaration
	// Only available inside functions
	fast := 25

	fmt.Println(x,y,a,b,fast,pi)
}
