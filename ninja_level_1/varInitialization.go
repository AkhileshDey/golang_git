package main

import (
	"fmt"
)

var x = 42
var y = `"james Bond"`
var z = true

func main() {
	s := fmt.Sprintln(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(s)
}
