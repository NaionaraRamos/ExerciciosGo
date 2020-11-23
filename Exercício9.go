package main

import (
	"fmt"
)

var x int = 200

func main() {
	//x := 200
	fmt.Printf("%d\n%#x\n%b\n", x, x, x)
	z := x << 1
	//fmt.Println(z)
	fmt.Printf("%d\n%#x\n%b", z, z, z)
}