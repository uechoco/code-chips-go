// https://talks.godoc.org/github.com/campoy/whispering-gophers/talk.slide#11
package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

const input = `A haiku is more
Than just a collection of
Well-formed syllables
`

func main() {
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
