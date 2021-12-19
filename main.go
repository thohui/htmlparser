package main

import (
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"github.com/thomas-huisman/htmlparser/parser"
)

func main() {
	bytes, err := ioutil.ReadFile("test.html")
	if err != nil {
		panic(err)
	}
	doc := parser.Parse(string(bytes))
	spew.Dump(doc.Root)
}
