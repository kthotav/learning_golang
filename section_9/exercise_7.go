package main

import "fmt"

func main() {
	a := 42
	if a < 24 {
		fmt.Println("Less than 24")
	} else if a == 42 {
		fmt.Println("Equal to 42")
	} else {
		fmt.Println("not equal to 42")
	}
}
