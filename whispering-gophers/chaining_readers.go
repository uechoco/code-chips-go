// https://talks.godoc.org/github.com/campoy/whispering-gophers/talk.slide#10
// https://talks.godoc.org/github.com/campoy/whispering-gophers/talk.slide#14
package main

import (
	"compress/gzip"
	"encoding/base64"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var r io.Reader
	r = strings.NewReader(data)
	r = base64.NewDecoder(base64.StdEncoding, r)
	r2, err := gzip.NewReader(r)
	//r2, err := gzip.NewReader(strings.NewReader(`This is not a gzip stream!!`))
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, r2)
}

const data = `H4sIAAAJbogA/1SOO5KDQAxE8zlFZ5tQXGCjjfYIjoURoPKgcY0E57f4VZlQXf2e+r8yOYbMZJhoZWRxz3wkCVjeReETS0VHz5fBCzpxxg/PbfrT/gacCjbjeiRNOChaVkA9RAdR8eVEw4vxa0Dcs3Fe2ZqowpeqG79L995l3VaMBUV/02OS+B6kMWikwG51c8n5GnEPr11F2/QJAAD//z9IppsHAQAA`
