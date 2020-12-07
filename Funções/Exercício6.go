package main

import (
	"fmt"
)

func main() {
	
	x := "Planetas..."
	func(x string) {
		fmt.Println(x)
	}(x)
}