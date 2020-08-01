package main

import "fmt"

type mynumber int
var x mynumber
func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)
}