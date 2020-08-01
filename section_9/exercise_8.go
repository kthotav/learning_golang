package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("should print")
	case false:
		fmt.Println("should not print")
	}
}
