package main

import "fmt"

func main() {
	// var fruitArray [2]string

	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"

	//Dec and assign
	// fruitArray := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArray)

	fruitSlice := []string{"Apple", "Orange", "Grape", "Pineapple"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
