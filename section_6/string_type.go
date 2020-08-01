package main

import "fmt"

func main() {
	const s = "Hello, Karthik!"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i])
	}

	for i, v := range s {
		fmt.Printf("At index %d: %c\n", i, v)
	}
}
