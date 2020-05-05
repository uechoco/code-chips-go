package main

import (
	"fmt"
)

func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func genInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	f := later()

	fmt.Println(f("Golang"))    // => ""
	fmt.Println(f("is"))        // => "Golang"
	fmt.Println(f("awesome!!")) // => "is"

	f2 := genInt()

	fmt.Println(f2()) // => "1"
	fmt.Println(f2()) // => "2"
	fmt.Println(f2()) // => "3"

	f3 := genInt()
	fmt.Println(f3()) // => "1"
	fmt.Println(f3()) // => "2"
}
