package main

import "fmt"

func main() {
	// var name string = "Gautham"
	var age int = 23
	const iscool = true

	//shorthand
	// name := "Gautham"
	size := 1.3

	name, email := "Gautham", "gautham@gmail.com"

	fmt.Println(name, age, iscool, email)
	fmt.Printf("%T\n", size)
}
