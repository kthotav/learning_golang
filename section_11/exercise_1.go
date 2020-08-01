package main

import "fmt"

func main() {
	a := [7]int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", a)
}
