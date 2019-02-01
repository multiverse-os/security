package main

import (
	"fmt"

	. "github.com/lostinblue/shodan"
)

func main() {
	shodan := New()
	ips, err := shodan.Search("Geth")
	if err != nil {
		fmt.Println("[Error] Failed to search", err)
	}else{
		fmt.Println("Found [", len(ips), "]")
	}
}
