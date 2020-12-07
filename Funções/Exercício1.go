package main

import (
	"fmt"
)

func inteiro() int{
	return 10
}

func inteiroestring() (int, string) {
	return 10, "Olá, planetinha!"
}

func main() {
	fmt.Println(inteiro())
	fmt.Println(inteiroestring())
}