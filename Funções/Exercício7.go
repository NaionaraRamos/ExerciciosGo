package main

import (
	"fmt"
)

func main() {

	x := "Olá, planetinhas..."
	a := func(x string) string {
		return x
	}
	fmt.Println(a(x))
	
	z := func(){
		fmt.Println("Olá, planetinhas...")
	}
	z()
}