package main

import "fmt"

type pizza int

var x pizza
var y int

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("x: %v\n", x)
	y = int(x)
	fmt.Printf("y: %v\n", y)
}
