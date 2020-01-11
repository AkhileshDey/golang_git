package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "James, Bond"
	z := true
	fmt.Printf("Print in a single line:- %v\t%v\t%v\n", x, y, z)
	fmt.Println("print in a multiple line:-")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
