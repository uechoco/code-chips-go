package main

import (
	"fmt"
)

func main() {
	// ----------------------------------------------
	// definitions
	// ----------------------------------------------

	// define a variable explicitly
	var a = 1
	fmt.Printf("%d", a)
	// define a variable implicitly
	b := 2
	fmt.Printf("%d", b)
	// define multiple variables
	var (
		s1 = "aaaa"
		c  = 2
	)
	fmt.Printf("%s %d", s1, c)

	// ----------------------------------------------
	// primitive types
	// ----------------------------------------------

	// boolean
	var b1 bool
	b1 = true
	b2 := false
	fmt.Printf("%v %v", b1, b2)

	// integers
	var (
		i1 int8    = 1
		i2 int16   = 1
		i3 int32   = 1
		i4 int64   = 1
		i5         = /*int*/ 1
		u1 uint8   = 1
		u2 uint16  = 1
		u3 uint32  = 1
		u4 uint64  = 1
		u5 uint    = 1
		u6 uintptr = 1
	)
	fmt.Printf("%v%v%v%v%v%v%v%v%v%v%v", i1, i2, i3, i4, i5, u1, u2, u3, u4, u5, u6)
	// integer cast to each other
	i4 = int64(i2)
	// wrap around
	n2 := 256
	b5 := byte(n2) // wrap around to 0
	fmt.Printf("%v", b5)
	// rune
	r := '松'
	fmt.Printf("%v", r)
	// string
	s := "Goの文字列"
	fmt.Printf("%s", s)
	// raw string literal
	s2 := `
Go String literal
Is raw string
	`
	fmt.Printf("%v", s2)
	// array
	a1 := [...]int{1, 2, 3}
	fmt.Printf("%v", a1)

}
