package main

import (
	"fmt"
)

type example int

var x example

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v", x)
}
