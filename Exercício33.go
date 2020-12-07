package main

import (
	"fmt"
)

type pessoa struct {
	nome            string
	sobrenome       string
	sabores_sorvete []string
}

func main() {

	pessoa1 := pessoa {
		nome: "Pessoa",
		sobrenome: "Um",
		sabores_sorvete: []string{"Flocos"}}
		
	pessoa2 := pessoa {"Pessoa", "Dois", []string{"Flocos"}}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sabores_sorvete {
		fmt.Println("\t", v)}
		
	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.sabores_sorvete {
		fmt.Println("\t", v)}
}
