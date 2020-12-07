package main

import (
	"fmt"
)

func main() {

x := struct {
	nome string
	sabor string
	ondetem []string
	vaibemcom map[string]string
}{
	nome: "Pizza",
	sabor: "Quatro Queijos",
	ondetem: []string{"Pizzaria", "Padaria"},
	vaibemcom: map[string]string{
		"no jantar": "Refrigerante",
		"no lanche": "Suco",
		},
}

	fmt.Println(x)
}