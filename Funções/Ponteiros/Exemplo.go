package main

import (
	"fmt"
)

func main() {

	x := 10
	//a := &x
	//fmt.Println(*a)
	//fmt.Println(*&x)
	//fmt.Printf("%T, %T", x, a)
	
	estarecebeumvalor(x)
	estarecebeumponteiro(&x)
}

func estarecebeumvalor(x int){
	x++
	fmt.Println(x)
}
	
func estarecebeumponteiro(x *int){
	*x++
	fmt.Println(*x)
}
