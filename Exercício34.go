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

	mapa := make(map[string]pessoa)

	mapa2 := map[string]pessoa{
		"Um": pessoa{
			nome:            "Pessoa",
			sobrenome:       "Um",
			sabores_sorvete: []string{"Flocos", "Abacaxi"}},

		"Dois": pessoa{"Pessoa", "Dois", []string{"Flocos", "Laranja"}},
	}

	mapa["Um"] = pessoa{
		nome:            "Pessoa",
		sobrenome:       "Um",
		sabores_sorvete: []string{"Flocos"}}

	mapa["Dois"] = pessoa{"Pessoa", "Dois", []string{"Flocos"}}

	for _, valor := range mapa {
		fmt.Println("Eu sou a pessoa", valor.nome, valor.sobrenome, "e meu sorvete favorito é", valor.sabores_sorvete[0])
	}
	fmt.Println()

	for _, valor := range mapa2 {
		fmt.Println("Eu sou a pessoa", valor.nome, valor.sobrenome, "e meus sorvetes favoritos são:")
		for _, valor := range valor.sabores_sorvete {
			fmt.Println("-", valor)
		}
		fmt.Println()
	}
}