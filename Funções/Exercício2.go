package main

import (
	"fmt"
)

func main() {
	s := []int{10, 20, 15, 30, 25}
	fmt.Println(funcao1(s...))
	fmt.Println(funcao2(s))
}

func funcao1(x ...int) int {
	soma := 0
	for _, valor := range x {
		soma += valor
	}
	return soma
}

func funcao2(x []int) int {
	soma := 0
	for _, valor := range x {
		soma += valor
	}
	return soma
}