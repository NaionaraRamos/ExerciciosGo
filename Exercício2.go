package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%v\n%v\n%v\n", x, y, z)
}