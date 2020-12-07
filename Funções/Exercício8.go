package main

import (
	"fmt"
)

func main() {

	x := retornaumafuncao()
	x()
}

func retornaumafuncao() func() {
	return func()  {
		fmt.Println("Hello, playground")
	}
}