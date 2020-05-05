package main

import (
	"fmt"
)

func main() {
	const (
		Pi = 3.1415926535897932384626433832795028841971693
	)

	f32 := float32(Pi)
	f64 := float64(Pi)

	fmt.Printf("%v", f32)
	fmt.Printf("%v", f64)

	const (
		A = iota // A == 0
		B        // B == 1
		C        // C == 2
	)

	fmt.Printf("%v %v %v", A, B, C)

	const (
		D = iota + 1 // D == 1
		E            // E == 2 (explicitly repeat `iota + 1`)
		F            // F == 3 (explicitly repeat `iota + 1`)
	)

	fmt.Printf("%v %v %v", D, E, F)
}
