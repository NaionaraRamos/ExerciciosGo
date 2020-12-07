package main

import (
	"fmt"
)

func main() {
	ss := [][]string{
		[]string{
			"Nai",
			"Ramos",
			"Ler",
		},
		[]string{
			"Ilan",
			"Vítor",
			"Jogar",
		},
		[]string{
			"Roni",
			"Ramos",
			"Sair com os amigos",
		},
	}
	fmt.Println(ss)
	
	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}
	}
}