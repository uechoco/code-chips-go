package main

import (
	"flag"
	"fmt"
)

var (
	// accept:
	//   -message 'OK'
	//   -message='OK'
	//   --message 'OK'
	//   --message='OK'
	message = flag.String("message", "", "what to say")
)

func main() {
	flag.Parse()
	fmt.Println(*message)
	fmt.Printf("processed argc:%d, remaining argc:%d\n", flag.NFlag(), flag.NArg())
}
