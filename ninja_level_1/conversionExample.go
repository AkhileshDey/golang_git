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
	fmt.Printf("%v\n", x)
	Y := int(x)
	fmt.Printf("%T\n", Y)
	fmt.Println(Y)
}
