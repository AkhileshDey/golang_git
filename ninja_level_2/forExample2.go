package main

import (
	"fmt"
)

func main() {
	birthdate := 1986
	for {
		fmt.Println(birthdate)
		if birthdate == 2020 {
			break
		}
		birthdate++
	}
}
