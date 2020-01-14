package main

import (
	"fmt"
)

func main() {
	birthdate := 1986
	for birthdate <= 2020 {
		fmt.Println(birthdate)
		birthdate++
	}
}
