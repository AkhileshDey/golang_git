package main

import (
	"fmt"
)

func main() {
	num1 := 10
	num2 := 10
	if num1 == num2 {
		fmt.Println("Its a match")
	} else if num1 != num2 {
		fmt.Println("not matched")
	} else {
		fmt.Println("This will not execute")
	}
}
