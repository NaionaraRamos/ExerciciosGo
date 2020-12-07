package main

import (
	"fmt"
)

func main() {

	//FUNÇÕES ANÔNIMAS

	x := 10

	func(x int) {
		fmt.Println(x, "x 10 =", x * 10)
	}(x)

	//FUNÇÃO COMO EXPRESSÃO
	x := 10

	a := func(x int) int {
		return x * 10
	}

	fmt.Println(x, "x 10 =", a(x))

	//RETORNANDO UMA FUNÇÃO
	x := retornaumafuncao()
	y := x(3)
	fmt.Println(y)

	fmt.Println(retornaumafuncao()(5))
}

func retornaumafuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
