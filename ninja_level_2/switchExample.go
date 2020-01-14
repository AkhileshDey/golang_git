package main

import (
	"fmt"
)

func main() {
	name := "James Bond"
	switch {
	default:
		fmt.Println("Not a Bond fan")
	case name == "James":
		fmt.Println("This is not James Bond")
	case name == "James Bond":
		fmt.Println("This is Bond, James Bond")

	}
}
