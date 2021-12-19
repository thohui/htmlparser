package parser

import (
	"strings"

	"golang.org/x/net/html"
)

type Element struct {
	Type       string
	Value      string
	Attributes map[string]string
	Children   []*Element
	Parent     *Element
}

type Document struct {
	Root *Element
}

func Parse(input string) *Document {
	doc := &Document{newElement()}

	tokens := html.NewTokenizer(strings.NewReader(input))
	var currentElement *Element = doc.Root
loop:
	for {
		current := tokens.Next()
		t := tokens.Token()
		switch current {
		case html.ErrorToken:
			break loop
		case html.StartTagToken:
			newElement := newElement()
			newElement.Type = t.Data
			newElement.Parent = currentElement
			currentElement.Children = append(currentElement.Children, newElement)
			newElement.Attributes = parseAttributes(t.Attr)
			currentElement = newElement
		case html.EndTagToken:
			if currentElement.Parent != nil {
				currentElement = currentElement.Parent
			}
		case html.TextToken:
			if t.Data != "" {
				currentElement.Value = t.Data
			}
		}
	}
	return doc
}

func parseAttributes(attributes []html.Attribute) map[string]string {
	m := make(map[string]string)
	for i := 0; i < len(attributes); i++ {
		attr := attributes[i]
		m[attr.Key] = attr.Val
	}
	return m
}

func newElement() *Element {
	return &Element{"", "", make(map[string]string), make([]*Element, 0), nil}
}
