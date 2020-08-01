package main

import "fmt"

func main() {
	a := (42 == 42)
	b := (42 <= 23)
	c := (26 >= 30)
	d := (32 != 32)
	e := (39 < 40)
	f := (68 > 20)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
