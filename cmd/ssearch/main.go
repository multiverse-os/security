package main

import (
	"fmt"

	. "github.com/multiverse-os/security-toolkit/shodan"
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
