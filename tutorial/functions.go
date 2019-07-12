package main

import "fmt"

func add(x int, y int) int {
	return x+y
}

// Naked return
func swap(x,y int)(x2,y2 int) {
	x2 = y
	y2 = x
	return
}

func main() {
	fmt.Println(add(42,13))
	fmt.Println(swap(50,20))
}
