package main

import (
	"fmt"
)

func main() {
	// ----------------------------------------------
	// if statement
	// ----------------------------------------------

	// basic
	x := 1
	if x == 1 {
		fmt.Println("1")
	} else if x == 2 {
		fmt.Println("2")
	} else if x == 3 {
		fmt.Println("3")
	} else {
		fmt.Println("other")
	}

	// if with a short statement (good for locality)
	if x, y := 1, 2; x < y {
		fmt.Println("x < y")
	}
	/*
		if _, err := doSomething(); err != nil {
		}
	*/

	// ----------------------------------------------
	// for statement
	// ----------------------------------------------

	// inifinte loop
	for {
		fmt.Println("inifinte loop")
		break
	}

	// with a condition
	n := 1
	for n < 10 {
		n++
	}

	// init and post statements (traditional for statement)
	for i := 1; i < 10; i++ {
		if i == 3 {
			continue
		}
	}

	// ranged for
	fruits := [...]string{"Apple", "Orange", "Banana"}
	for i, s := range fruits {
		fmt.Printf("index:=%d, name:=%s", i, s)
	}

	// ----------------------------------------------
	// switch statement
	// ----------------------------------------------
	switch x {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	case 5:
		fmt.Println("5")
		fallthrough
	default:
		fmt.Println("more than or equal to 5")
	}

	switch {
	case n > 0 && n < 3:
		fmt.Println("0 < n < 3")
	case n > 3 && n < 6:
		fmt.Println("3 < n < 6")
	}

	var s3 interface{}
	s3 = "hoge"
	switch s3.(type) {
	case bool:
		fmt.Println("bool")
	case string:
		fmt.Println("string")
	}

	// ----------------------------------------------
	// defer statement
	// ----------------------------------------------
	runDefer := func() {
		defer fmt.Println("defer")
		fmt.Println("done")
	}
	runDefer()

	runDefer2 := func() {
		defer fmt.Println("defer1")
		defer fmt.Println("defer2")
		defer fmt.Println("defer3")
		fmt.Println("done")
	}
	runDefer2()

	// ----------------------------------------------
	// panic and recover statement
	// ----------------------------------------------
	func1 := func() {
		defer func() {
			if x := recover(); x != nil {
				/* x は panic の引数 interface{} */
				fmt.Println(x)
				fmt.Println("and recovered")
			}
		}()
		panic("paniced")
		fmt.Println("This statement is not run")
	}
	func1()
}
