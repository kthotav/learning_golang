package main

import "fmt"

func main() {
	favSport := "basketball"
	switch favSport {
	case "baseball":
		fmt.Println("Baseball is your favorite sport!")
	case "basketball":
		fmt.Println("Basketbal is your favorite sport!")
	case "soccer":
		fmt.Println("Soccer is your favorite sport!")
	}
}
