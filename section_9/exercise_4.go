package main

import "fmt"

func main() {
	bd := 1994
	for {
		if bd > 2021 {
			break
		} else {
			fmt.Println(bd)
			bd++
		}
	}
}
