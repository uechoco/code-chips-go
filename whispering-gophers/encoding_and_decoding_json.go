package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Site struct {
	Title string
	URL   string
}

var sites = []Site{
	{"The Go Programming Language", "http://golang.org"},
	{"Google", "http://google.com"},
}

const stream = `
{"Title": "The Go Programming Language", "URL": "http://golang.org"}
{"Title": "Google", "URL": "http://google.com"}
`

func main() {
	enc := json.NewEncoder(os.Stdout)
	for _, s := range sites {
		err := enc.Encode(s)
		if err != nil {
			log.Fatal(err)
		}
	}

	dec := json.NewDecoder(strings.NewReader(stream))
	for dec.More() {
		var s Site
		if err := dec.Decode(&s); err != nil {
			log.Fatal(err)
		}
		fmt.Println(s.Title, s.URL)
	}
}
