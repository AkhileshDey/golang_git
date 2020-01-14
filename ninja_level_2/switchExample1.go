package main

import (
	"fmt"
)

func main() {
	name := "James Bond"
	switch name {
	default:
		fmt.Println("Not a Bond fan")
	case "James":
		fmt.Println("This is not James Bond")
	case "James Bond":
		fmt.Println("This is Bond, James Bond")

	}
}
