package main

import (
	"fmt"
)

//RECURSIVIDADE
func main() {

	fmt.Println(fatorial(5))
	fmt.Println(loops(5))
}

func fatorial(x int) int {

	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}

func loops(x int) int {
	total := 1
	for x > 1 {
		total *= x
		x--
	}
	return total
}