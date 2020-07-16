package main

import "fmt"

const name, age = "Ahsoka Tano", 22

func getNameAndAge(name string, age int) string {
	return fmt.Sprintf("Hello, name name is %s and I am %d years old!", name, age)
}
func main() {
	nameAndAge := getNameAndAge(name, age)
	fmt.Println(nameAndAge)
}
