package main

import (
	"fmt"

	"github.com/akakou/crtsh"
)

func main() {
	data, err := crtsh.Fetch("test.ochano.co", "expired")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Data: %v\n", data)
}
