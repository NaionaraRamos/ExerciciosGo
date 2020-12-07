package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	fmt.Println(x[:3])
	fmt.Println(x[4:])
	fmt.Println(x[1:7])
	fmt.Println(x[2:9])
	fmt.Println(x[2:len(x)-1])
}