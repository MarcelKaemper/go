package main

import "fmt"

func main() {
	x := 0
	for i := 0; i<10; i++ {
		x += i
	}

	fmt.Println(x)
}
