package main

import (
	"fmt"
)

func main() {
	// create simple slice
	s1 := []int{1, 2, 3}
	for i, v := range s1 {
		fmt.Printf("s1[%d] == %v\n", i, v)
	}
	// get length of the slice
	fmt.Println("s1 length is ", len(s1))
	// get capacity of the slice
	fmt.Println("s1 capacity is ", cap(s1))

	// create slice via make embedded function
	s2 := make([]int, 3)
	for i, v := range s2 {
		fmt.Printf("s2[%d] == %v\n", i, v)
	}
	// get length of the slice
	fmt.Println("s2 length is ", len(s2))
	// get capacity of the slice
	fmt.Println("s2 capacity is ", cap(s2))

	// create slice with capacity hint via make embedded function
	s3 := make([]int, 3, 8)
	for i, v := range s3 {
		fmt.Printf("s3[%d] == %v\n", i, v)
	}
	// get length of the slice
	fmt.Println("s3 length is ", len(s3))
	// get capacity of the slice
	fmt.Println("s3 capacity is ", cap(s3))

	s4 := make([]int, 3)
	s4[1] = 3
	fmt.Printf("s4 pointer:%p len:%d, cap:%d, values:%v\n", s4, len(s4), cap(s4), s4)
	s4 = append(s4, 4)
	fmt.Printf("s4 pointer:%p len:%d, cap:%d, values:%v\n", s4, len(s4), cap(s4), s4)
}
