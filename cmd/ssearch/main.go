package main

import (
	. "github.com/lostinblue/shodan"
)

func main() {
	shodan := New()
	shodan.Search("Geth")
}
