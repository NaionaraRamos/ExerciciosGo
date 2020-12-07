package main

import (
	"fmt"
)

func main() {
	fmt.Println(soma(10, 6))
	basica()
	cumprimento("tarde")
	fmt.Println(soma_e_comprimento(1, 2, 3, 4, 5, 6))
}

func soma(a, b int) int {
	return a + b
}

func soma_e_comprimento(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)
}

func basica() {
	fmt.Println("Planetinhas...")
}

func cumprimento(s string) {
	if s == "manhã" {
		fmt.Println("Bom dia!")
	} else if s == "tarde" {
		fmt.Println("Boa tarde!")
	} else {
		fmt.Println("Boa noite!")
	}
}
